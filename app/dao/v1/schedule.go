package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
	"time"
)

type scheduleDAO struct{}

var ScheduleDAO *scheduleDAO = &scheduleDAO{}

func (s *scheduleDAO) DoGetWithinDaysScheduleList(startDate models.Date, generateDays int) ([]models.Schedule, error) {
	// var schedules []*models.Schedule
	schedules := []models.Schedule{}
	date := time.Time(startDate)
	if err := core.DB.Where("date >= ? AND date < ?", date, date.AddDate(0, 0, generateDays)).Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *scheduleDAO) DoCreateSchedule(schedule []*models.Schedule) error {
	if len(schedule) == 0 {
		return nil
	}
	if err := core.DB.Create(schedule).Error; err != nil {
		return err
	}
	return nil
}

func (s *scheduleDAO) DoGetSchedule(gsi *models.GetScheduleIn) (*models.Schedule, error) {
	schedule := &models.Schedule{}
	if err := core.DB.Find(schedule, &models.Schedule{
		Date:      gsi.Date,
		UserID:    gsi.UserID,
		ServiceID: gsi.ServiceID,
	}).Error; err != nil {
		return nil, err
	}
	return schedule, nil
}
