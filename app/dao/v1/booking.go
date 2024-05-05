package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
)

type bookingDAO struct {
}

var BookingDAO *bookingDAO = &bookingDAO{}

func (s *bookingDAO) DoCreateBooking(cbi *models.CreateBookingIn) {
	core.DB.Create(&models.Booking{})
}
