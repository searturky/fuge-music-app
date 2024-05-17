package service_v1

import (
	"fuge/app/constant"
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
	"fuge/app/utils"
	"time"
)

type scheduleService struct{}

var ScheduleService *scheduleService = &scheduleService{}

func (s *scheduleService) QuickGenerate(qgi *models.QuickGenerateIn) error {
	// 先查询是否已经有排班
	schedules, err := daos.ScheduleDAO.DoGetWithinDaysScheduleList(qgi.StartDate, qgi.GenerateDays)
	if err != nil {
		return err
	}
	// 已经生成过的排班
	scheduledDays := make(map[string]struct{})
	for _, schedule := range schedules {
		dateStr := schedule.Date.Format(time.DateOnly)
		scheduledDays[dateStr] = struct{}{}
	}
	// 期望生成的排班
	expectedDays := make(map[string]struct{})
	for i := 0; i < qgi.GenerateDays; i++ {
		date := qgi.StartDate.Time().AddDate(0, 0, i)
		timeStr := date.Format(time.DateOnly)
		expectedDays[timeStr] = struct{}{}
	}
	// 需要生成的排班，即期望生成的排班中没有的
	shouldGenerateDays := []string{}
	for day := range expectedDays {
		if _, ok := scheduledDays[day]; !ok {
			shouldGenerateDays = append(shouldGenerateDays, day)
		}
	}
	if err := generateSchedule(qgi, shouldGenerateDays); err != nil {
		return err
	}
	return nil
}

func generateSchedule(qgi *models.QuickGenerateIn, shouldGenerateDays []string) error {
	service, err := daos.ServiceDAO.DoGetServiceByID(qgi.ServiceID)
	if err != nil {
		return err
	}
	schedules := []*models.Schedule{}
	for _, day := range shouldGenerateDays {
		// 生成排班
		scheduleTime, err := time.ParseInLocation(time.DateOnly, day, time.Local)
		if err != nil {
			return err
		}
		slots, err := generateTimeSlots(
			qgi.DailyStartTime,
			qgi.DailyEndTime,
			service.TimePeriod,
			scheduleTime,
		)
		if err != nil {
			return err
		}
		schedule := &models.Schedule{
			UserID:         qgi.UserID,
			ServiceID:      service.ID,
			Date:           scheduleTime,
			DailyStartTime: qgi.DailyStartTime,
			DailyEndTime:   qgi.DailyEndTime,
			TimeSlots:      slots,
			TimePeriod:     service.TimePeriod,
		}
		schedules = append(schedules, schedule)
	}
	return daos.ScheduleDAO.DoCreateSchedule(schedules)
}

func generateTimeSlots(startTime, endTime string, timePeriod int, scheduleDate time.Time) ([]string, error) {
	stascheduleDay := scheduleDate.Format(time.DateOnly)
	start, err := time.ParseInLocation(constant.DateTimeNoSecond, stascheduleDay+" "+startTime, time.Local)
	if err != nil {
		return nil, err
	}
	end, err := time.ParseInLocation(constant.DateTimeNoSecond, stascheduleDay+" "+endTime, time.Local)
	if err != nil {
		return nil, err
	}
	timeSlots := []string{}
	for start.Before(end) {
		startStr := start.Format(constant.TimeNoSecond)
		timeSlots = append(timeSlots, startStr)
		start = start.Add(time.Duration(timePeriod) * time.Minute)
	}
	return timeSlots, nil
}

func (s *scheduleService) GetScheduleByUserAndDate(gsi *models.GetScheduleIn) (*models.GetScheduleOut, error) {
	schedule, err := daos.ScheduleDAO.DoGetSchedule(gsi)
	if err != nil {
		return nil, err
	}
	slots, err := getScheduleTimeSlotsShowWithState(schedule)
	if err != nil {
		return nil, err
	}
	gso := &models.GetScheduleOut{
		ScheduleID:        schedule.ID,
		Date:              schedule.Date.Format(time.DateOnly),
		TimePeriod:        schedule.TimePeriod,
		StatefulTimeSlots: slots,
	}
	return gso, nil
}

func getScheduleTimeSlotsShowWithState(schedule *models.Schedule) ([]models.StatefulTimeSlot, error) {
	/*
		返回当日含预约状态的排班时间段
	*/
	statefulTimes, err := buildStateTimeSlots(schedule)
	if err != nil {
		return nil, err
	}
	return statefulTimes, nil
}

func buildStateTimeSlots(schedule *models.Schedule) ([]models.StatefulTimeSlot, error) {
	/*
		根据预约时间段，构建排班时间段的预约状态
	*/
	statefulTimeSlots := []models.StatefulTimeSlot{}
	// 获取该当日门店房间所有预约
	roomBookingMap, err := getRoomBookingMap(schedule)
	if err != nil {
		return nil, err
	}

	for _, slot := range schedule.TimeSlots {
		isAvailable, err := slotIsAvailable(slot, schedule, roomBookingMap)
		if err != nil {
			return nil, err
		}
		statefulTimeSlots = append(statefulTimeSlots, models.StatefulTimeSlot{
			TimeSlot:    slot,
			IsAvailable: isAvailable,
		})
	}
	return statefulTimeSlots, nil
}

func slotIsAvailable(slot string, schedule *models.Schedule, args ...any) (bool, error) {
	// 获取服务者当日已预约时间段
	serverBookings, err := daos.BookingDAO.DoGetBookingsByUserIDAndDate(schedule.UserID, schedule.Date)
	if err != nil {
		return false, err
	}
	// 构造服务者的已预约时间区间
	serverBookedTimes := []time.Time{}
	date := schedule.Date
	for _, booking := range serverBookings {
		startDatetimeStr := booking.Date.Format(time.DateOnly) + " " + booking.BookingTime
		startTime, _ := time.ParseInLocation(constant.DateTimeNoSecond, startDatetimeStr, time.Local)
		serverBookedTimes = append(serverBookedTimes, startTime)
		endTime := startTime.Add(time.Duration(booking.BookingTimePeriod) * time.Minute)
		serverBookedTimes = append(serverBookedTimes, endTime)
	}
	slotDatetimeStr := date.Format(time.DateOnly) + " " + slot
	slotTime, err := time.ParseInLocation(constant.DateTimeNoSecond, slotDatetimeStr, time.Local)
	if err != nil {
		return false, err
	}
	// 获取该当日门店房间所有预约
	var roomBookingMap map[int]*utils.BitMap
	isBitMapInArgs := false
	if len(args) > 0 {
		if value, ok := args[0].(map[int]*utils.BitMap); ok {
			roomBookingMap = value
			isBitMapInArgs = true
		}
	}
	if !isBitMapInArgs {
		roomBookingMap, err = getRoomBookingMap(schedule)
		if err != nil {
			return false, err
		}
	}

	// 计算服务者的可用时间
	for i := 0; i < len(serverBookedTimes); i += 2 {
		// 该服务者该时间段已经没有时间
		if slotTime.Compare(serverBookedTimes[i]) >= 0 && slotTime.Compare(serverBookedTimes[i+1]) <= 0 {
			return false, nil
		}
	}
	// 计算可用房间
	availableRooms, err := getSlotAvailableRooms(slot, schedule, roomBookingMap)
	if err != nil || len(availableRooms) == 0 {
		return false, err
	}
	return true, nil
}

func getRoomBookingMap(schedule *models.Schedule) (map[int]*utils.BitMap, error) {
	service, err := daos.ServiceDAO.DoGetServiceByID(schedule.ServiceID)
	if err != nil {
		return nil, err
	}
	RoomIDs := []int{}
	for _, room := range service.Rooms {
		RoomIDs = append(RoomIDs, room.ID)
	}
	allStoreBooking, err := daos.BookingDAO.DoGetAllBookings(service.StoreID, schedule.Date, RoomIDs)
	if err != nil {
		return nil, err
	}
	// 构造房间预约时间段Map位图
	roomMap := make(map[int]*utils.BitMap)
	for _, booking := range allStoreBooking {
		if _, ok := roomMap[booking.RoomID]; !ok {
			roomMap[booking.RoomID] = utils.NewBitMap()
		}
		startMinute, err := utils.GetMinuteFromTimeStr(booking.BookingTime)
		if err != nil {
			return nil, err
		}
		endMinute := startMinute + booking.BookingTimePeriod
		for i := startMinute; i <= endMinute; i++ {
			roomMap[booking.RoomID].Set(uint(i))
		}
	}
	return roomMap, nil
}

func getSlotAvailableRooms(slot string, schedule *models.Schedule, args ...any) ([]int, error) {
	var roomBookingMap map[int]*utils.BitMap
	isBitMapInArgs := false
	var err error
	if len(args) > 0 {
		if value, ok := args[0].(map[int]*utils.BitMap); ok {
			roomBookingMap = value
			isBitMapInArgs = true
		}
	}
	if !isBitMapInArgs {
		roomBookingMap, err = getRoomBookingMap(schedule)
		if err != nil {
			return nil, err
		}
	}
	availableRooms := []int{}
	slotStartMinute, err := utils.GetMinuteFromTimeStr(slot)
	if err != nil {
		return nil, err
	}
	slotEndMinute := slotStartMinute + schedule.TimePeriod
	for roomID, bitmap := range roomBookingMap {
		isAvailable := true
		for i := slotStartMinute; i <= slotEndMinute; i++ {
			if bitmap.Check(uint(i)) {
				isAvailable = false
				break
			}
		}
		if isAvailable {
			availableRooms = append(availableRooms, roomID)
		}
	}

	return availableRooms, nil
}
