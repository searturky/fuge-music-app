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

func (s *bookingDAO) DoGetBookingsBySchedule(schedule *models.Schedule) ([]*models.Booking, error) {
	userID := schedule.UserID
	bookings := []*models.Booking{}
	if err := core.DB.Where(
		"user_id = ?", userID,
	).Where(
		"date = ?", schedule.Date,
	).Where(
		"is_cancel = ?", false,
	).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
