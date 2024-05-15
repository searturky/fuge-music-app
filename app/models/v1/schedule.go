package models_v1

import (
	"time"

	"github.com/lib/pq"
)

type Schedule struct {
	BaseModel
	// StoreID           int         `gorm:"type:int; default:null; comment:门店ID"`
	// RoomID            int         `gorm:"type:int; default:null; comment:房间ID"`
	// ServiceCategoryID int         `gorm:"type:int; default:null; comment:服务分类ID"`
	UserID         int            `gorm:"type:int; default:null; comment:服务提供者ID"`
	ServiceID      int            `gorm:"type:int; default:null; comment:服务ID"`
	Date           time.Time      `gorm:"type:timestamptz(6); comment:日期"`
	DailyStartTime string         `gorm:"type:varchar(30); comment:每日开始时间"`
	DailyEndTime   string         `gorm:"type:varchar(30); comment:每日结束时间"`
	TimePeriod     int            `gorm:"type:int; default:1; check:time_period > 0 and time_period < 1440 and 3600 % time_period = 0; comment:时间段"`
	TimeSlots      pq.StringArray `gorm:"type:varchar(12)[]; comment:每日时间段"`
}

type QuickGenerateIn struct {
	UserID         int    `json:"user_id" binding:"required" example:"1"`
	ServiceID      int    `json:"service_id" binding:"required" example:"1"`
	StartDate      Date   `json:"start_date" binding:"required" example:"2024-05-15"`
	GenerateDays   int    `json:"generate_days" binding:"required" example:"7"`
	DailyStartTime string `json:"daily_start_time" binding:"required" example:"09:00"`
	DailyEndTime   string `json:"daily_end_time" binding:"required" example:"21:00"`
}

type GetScheduleIn struct {
	ServiceID int       `form:"s" binding:"required" example:"1" description:"服务ID"`
	UserID    int       `form:"u" binding:"required" example:"1" description:"服务者用户ID"`
	Date      time.Time `form:"d" binding:"required" example:"2024-05-15" description:"日期2024-05-15" time_format:"2006-01-02"`
}
