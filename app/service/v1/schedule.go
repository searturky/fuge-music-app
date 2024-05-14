package service_v1

import (
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
		dateStr := schedule.Date.Format("2006-01-02")
		scheduledDays[dateStr] = struct{}{}
	}
	// 期望生成的排班
	expectedDays := make(map[string]struct{})
	for i := 0; i < qgi.GenerateDays; i++ {
		date := qgi.StartDate.Time().AddDate(0, 0, i)
		timeStr := date.Format("2006-01-02")
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
		scheduleTime, err := time.ParseInLocation("2006-01-02", day, time.Local)
		if err != nil {
			return err
		}
		slots, err := generateTimeSlots(
			service.DailyStartTime,
			service.DailyEndTime,
			service.TimePeriod,
			scheduleTime,
		)
		if err != nil {
			return err
		}
		schedule := &models.Schedule{
			ServiceID:      service.ID,
			Date:           scheduleTime,
			DailyStartTime: service.DailyStartTime,
			DailyEndTime:   service.DailyEndTime,
			TimeSlots:      slots,
		}
		schedules = append(schedules, schedule)
	}
	return daos.ScheduleDAO.DoCreateSchedule(schedules)
}

func generateTimeSlots(startTime, endTime string, timePeriod int, scheduleDate time.Time) ([]string, error) {
	stascheduleDay := scheduleDate.Format("2006-01-02")
	start, err := time.ParseInLocation("2006-01-02 15:04", stascheduleDay+" "+startTime, time.Local)
	if err != nil {
		return nil, err
	}
	end, err := time.ParseInLocation("2006-01-02 15:04", stascheduleDay+" "+endTime, time.Local)
	if err != nil {
		return nil, err
	}
	timeSlots := []string{}
	for start.Before(end) {
		startStr := start.Format("15:04")
		timeSlots = append(timeSlots, startStr)
		start = start.Add(time.Duration(timePeriod) * time.Minute)
	}
	return timeSlots, nil
}
