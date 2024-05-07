package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
	"time"
)

type scheduleDAO struct{}

var ScheduleDAO *scheduleDAO = &scheduleDAO{}

func (s *scheduleDAO) DoGetWithinDaysScheduleList(days int) []*models.Schedule {
	// var schedules []*models.Schedule
	schedules := []*models.Schedule{}
	today := time.Now().Format("2006-01-02")
	core.DB.Where("date >= ? AND date <= ?", today, time.Now().AddDate(0, 0, days).Format("2006-01-02")).Find(schedules)
	return schedules
}
