package service_v1

import (
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
)

type serviceService struct {
}

var ServiceService *serviceService = &serviceService{}

func (s *serviceService) CreateService(csi *models.CreateServiceIn) {
	daos.ServiceDAO.DoCreateService(csi)
}
