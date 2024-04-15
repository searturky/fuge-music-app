package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type roomDAO struct {
}

var RoomDAO *roomDAO = &roomDAO{}

func (s *roomDAO) DoGetRoomsByStoreId(storeId int) ([]models.Room, error) {
	var rooms []models.Room
	err := core.DB.Where("store_id = ?", storeId).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
