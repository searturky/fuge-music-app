package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type roomDAO struct {
}

var RoomDAO *roomDAO = &roomDAO{}

func (s *roomDAO) DoGetRoomsByStoreId(storeId int) []models.Room {
	var rooms []models.Room
	core.DB.Where("store_id = ?", storeId).Find(&rooms)
	return rooms
}
