package service_v1

import (
	"fuge/app/constant"
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
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
	service := daos.ServiceDAO.DoGetServiceByID(qgi.ServiceID)
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

func (s *scheduleService) GetScheduleByUserAndDate(gsi *models.GetScheduleIn) ([]*models.Schedule, error) {
	schedule, err := daos.ScheduleDAO.DoGetSchedule(gsi)
	if err != nil {
		return nil, err
	}
	slots, err := getScheduleTimeSlotsShowWithState(schedule)
	if err != nil {
		return nil, err
	}
	print(slots)
	// return daos.ScheduleDAO.DoGetScheduleList(gsi.ServiceID, gsi.Date)
	return nil, nil
}

func getScheduleTimeSlotsShowWithState(schedule *models.Schedule) ([]string, error) {
	// 返回当日排班时间段，包含预约状态
	timeSlots := schedule.TimeSlots
	// 获取当日已预约时间段
	bookings, err := daos.BookingDAO.DoGetBookingsBySchedule(schedule)
	if err != nil {
		return nil, err
	}
	print(timeSlots, bookings)
	return nil, nil
}
