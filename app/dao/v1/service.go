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
