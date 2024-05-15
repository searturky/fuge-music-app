package models_v1

import (
	"time"
)

type Booking struct {
	BaseModel
	StoreID           int       `gorm:"required; not null; index; comment:门店ID"`
	RoomID            int       `gorm:"required; not null; index; comment:房间ID"`
	ServiceID         int       `gorm:"required; not null; index; comment:服务ID"`
	UserID            int       `gorm:"required; not null; index; comment:用户ID"`
	BookingUserID     int       `gorm:"required; not null; index; comment:预约用户ID"`
	Date              time.Time `gorm:"type:timestamptz(6); comment:日期 2021-01-01"`
	BookingTime       string    `gorm:"type:varchar(30); comment:预约时间 10:00"`
	BookingTimePeriod uint      `gorm:"required; not null; comment:预约时间间隔（分）"`
	IsSigned          bool      `gorm:"type:boolean; default:false; comment:是否已签到"`
	IsCancel          bool      `gorm:"type:boolean; default:false; comment:是否已取消"`
}

type CreateBookingIn struct {
	Date          Date      `json:"date" binding:"required" type:"string" format:"date"`
	StartDateTime time.Time `json:"startDateTime" binding:"required" type:"string" format:"date-time"`
	EndDateTime   time.Time `json:"endDateTime" binding:"required" type:"string" format:"date-time"`
}
