package dao_v1

import (
	"fuge/app/core"
	models "fuge/app/models/v1"
	"time"
)

type bookingDAO struct {
}

var BookingDAO *bookingDAO = &bookingDAO{}

func (s *bookingDAO) DoCreateBooking(cbi *models.CreateBookingIn) {
	core.DB.Create(&models.Booking{})
}

func (s *bookingDAO) DoGetBookingsByUserIDAndDate(userID int, date time.Time) ([]*models.Booking, error) {
	bookings := []*models.Booking{}
	if err := core.DB.Where(
		"user_id = ?", userID,
	).Where(
		"date = ?", date,
	).Where(
		"is_cancel = ?", false,
	).Order(
		"booking_time",
	).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *bookingDAO) DoGetAllBookings(storeID int, date time.Time, roomIDs []int) ([]*models.Booking, error) {
	bookings := []*models.Booking{}
	if err := core.DB.Where(
		"store_id = ?", storeID,
	).Where(
		"date = ?", date,
	).Where(
		"is_cancel = ?", false,
	).Where(
		"room_id IN (?)", roomIDs,
	).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
