package models_v1

import (
	"time"
)

type Booking struct {
	BaseModel
	StoreID           int    `gorm:"required; not null; index"`
	RoomID            int    `gorm:"required; not null; index"`
	ServiceID         int    `gorm:"required; not null; index"`
	UserID            int    `gorm:"required; not null; index"`
	BookUserID        int    `gorm:"required; not null; index"`
	Date              string `gorm:"type:varchar(30);"`
	BookingTime       string `gorm:"type:varchar(30);"`
	BookingTimePeriod uint   `gorm:"required; not null;"`
	IsSigned          bool   `gorm:"type:boolean; default:false"`
	IsCancel          bool   `gorm:"type:boolean; default:false"`
}

type CreateBookingIn struct {
	Date          Date      `json:"date" binding:"required" type:"string" format:"date"`
	StartDateTime time.Time `json:"startDateTime" binding:"required" type:"string" format:"date-time"`
	EndDateTime   time.Time `json:"endDateTime" binding:"required" type:"string" format:"date-time"`
}
