package models_v1

import "time"

type Schedule struct {
	BaseModel
	StoreID           int         `gorm:"type:int; default:null"`
	RoomID            int         `gorm:"type:int; default:null"`
	ServiceCategoryID int         `gorm:"type:int; default:null"`
	ServiceID         int         `gorm:"type:int; default:null"`
	UserID            int         `gorm:"type:int; default:null"`
	Date              time.Time   `gorm:"type:timestamptz(6)"`
	DailyStartTime    string      `gorm:"type:varchar(30);"`
	DailyEndTime      string      `gorm:"type:varchar(30);"`
	TimePeriod        int         `gorm:"type:int; default:1; check:time_period > 0 and time_period < 1440 and 3600 % time_period = 0"`
	TimeSlots         []time.Time `gorm:"type:timestamptz(6)[]"`
}

type QuickGenerateIn struct {
	StoreID           int  `json:"store_id" binding:"required"`
	RoomID            int  `json:"room_id" binding:"required"`
	ServiceCategoryID int  `json:"service_category_id"`
	ServiceID         int  `json:"service_id" binding:"required"`
	UserID            int  `json:"user_id" binding:"required"`
	Date              Date `json:"date" binding:"required"`
	Days              int  `json:"days" binding:"required"`
}
