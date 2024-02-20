package models_v1

import (
	"time"
)

type Booking struct {
	BaseModel
	UserID        uint
	Date          time.Time `gorm:"required; not null; index"`
	StartDateTime time.Time
	EndDateTime   time.Time
	LocationID    uint
}

type CreateBookingIn struct {
	Date          Date      `json:"date" binding:"required" type:"string" format:"date"`
	StartDateTime time.Time `json:"startDateTime" binding:"required" type:"string" format:"date-time"`
	EndDateTime   time.Time `json:"endDateTime" binding:"required" type:"string" format:"date-time"`
}
