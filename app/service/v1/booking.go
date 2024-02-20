package service_v1

import (
	daos "fuge/app/dao/v1"
	models "fuge/app/models/v1"
)

type bookingService struct {
}

var BookingService *bookingService = &bookingService{}

func (s *bookingService) CreateBooking(cbi *models.CreateBookingIn) {
	daos.BookingDAO.DoCreateBooking(cbi)
}
