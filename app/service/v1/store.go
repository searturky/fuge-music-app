package service_v1

import (
	daos "fuge/app/dao/v1"
)

type roomService struct {
}

var RoomService *roomService = &roomService{}

func (s *roomService) GetRoomByStoreId(storeId int) any {
	rooms := daos.RoomDAO.DoGetRoomsByStoreId(storeId)
	return rooms
}
