package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type serviceDAO struct {
}

var ServiceDAO *serviceDAO = &serviceDAO{}

func (s *serviceDAO) DoCreateService(csi *models.CreateServiceIn) {
	if err := core.DB.Create(&models.Service{
		Name:        csi.Name,
		Description: csi.Description,
	}).Error; err != nil {
		panic(err)
	}
}

func (s *serviceDAO) DoGetServiceByID(id int) (*models.Service, error) {
	service := &models.Service{}
	if err := core.DB.Where("id = ?", id).First(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}
