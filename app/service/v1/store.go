package service_v1

import (
	daos "fuge/app/dao/v1"
)

type roomService struct {
}

var RoomService *roomService = &roomService{}

func (s *roomService) GetRoomByStoreId(storeId int) (any, error) {
	rooms, err := daos.RoomDAO.DoGetRoomsByStoreId(storeId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}
