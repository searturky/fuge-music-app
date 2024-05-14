package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
	"time"
)

type scheduleDAO struct{}

var ScheduleDAO *scheduleDAO = &scheduleDAO{}

func (s *scheduleDAO) DoGetWithinDaysScheduleList(startDate models.Date, generateDays int, userId int) ([]models.Schedule, error) {
	// var schedules []*models.Schedule
	schedules := []models.Schedule{}
	date := time.Time(startDate)
	if err := core.DB.Where("date >= ? AND date < ?", date, date.AddDate(0, 0, generateDays)).Where("user_id = ?", userId).Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *scheduleDAO) DoCreateSchedule(schedule []*models.Schedule) error {
	if err := core.DB.Create(schedule).Error; err != nil {
		return err
	}
	return nil
}
