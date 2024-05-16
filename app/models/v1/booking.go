package models_v1

import (
	"time"
)

type Booking struct {
	BaseModel
	StoreID           int       `gorm:"required; not null; index; comment:门店ID"`
	RoomID            int       `gorm:"index; comment:房间ID"`
	ServiceID         int       `gorm:"required; not null; index; comment:服务ID"`
	UserID            int       `gorm:"required; not null; index; comment:用户ID"`
	BookingUserID     int       `gorm:"required; not null; index; comment:预约用户ID"`
	Date              time.Time `gorm:"type:timestamptz(6); comment:日期 2021-01-01"`
	BookingTime       string    `gorm:"type:varchar(30); comment:预约时间 10:00"`
	BookingTimePeriod int       `gorm:"required; not null; comment:预约时间间隔（分）"`
	IsSigned          bool      `gorm:"type:boolean; default:false; comment:是否已签到"`
	IsCancel          bool      `gorm:"type:boolean; default:false; comment:是否已取消"`
}

type CreateBookingIn struct {
	StoreID           int       `json:"store_id" binding:"required" example:"1"`
	RoomID            int       `json:"room_id" binding:"required" example:"1"`
	ServiceID         int       `json:"service_id" binding:"required" example:"1"`
	UserID            int       `json:"user_id" binding:"required" example:"1"`
	BookingUserID     int       `json:"booking_user_id" binding:"required" example:"1"`
	ScheduleID        int       `json:"schedule_id" binding:"required" example:"1"`
	Date              time.Time `json:"date" binding:"required" example:"2024-05-15" time_format:"2006-01-02"`
	BookingTime       string    `json:"booking_time" binding:"required" example:"09:00"`
	BookingTimePeriod int       `json:"booking_time_period" binding:"required" example:"60"`
}
