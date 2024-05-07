package service_v1

import (
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
)

type scheduleService struct{}

var ScheduleService *scheduleService = &scheduleService{}

func (s *scheduleService) QuickGenerate(qgi *models.QuickGenerateIn) {
	// 先查询是否已经有排班
	day := qgi.Days
	schedules := daos.ScheduleDAO.DoGetWithinDaysScheduleList(day)
	print(schedules)
	generateSchedule(qgi, "2024-5-7", "2024-5-14")
}

func generateSchedule(qgi *models.QuickGenerateIn, startDay string, endDay string) {
	service := daos.ServiceDAO.DoGetServiceByID(qgi.ServiceID)
	print(service)
}
