package models_v1

import (
	"github.com/lib/pq"
	"time"
)

type Schedule struct {
	BaseModel
	// StoreID           int         `gorm:"type:int; default:null; comment:门店ID"`
	// RoomID            int         `gorm:"type:int; default:null; comment:房间ID"`
	// ServiceCategoryID int         `gorm:"type:int; default:null; comment:服务分类ID"`
	ServiceID      int            `gorm:"type:int; default:null; comment:服务ID"`
	Date           time.Time      `gorm:"type:timestamptz(6); comment:日期"`
	DailyStartTime string         `gorm:"type:varchar(30); comment:每日开始时间"`
	DailyEndTime   string         `gorm:"type:varchar(30); comment:每日结束时间"`
	TimePeriod     int            `gorm:"type:int; default:1; check:time_period > 0 and time_period < 1440 and 3600 % time_period = 0; comment:时间段"`
	TimeSlots      pq.StringArray `gorm:"type:varchar(12)[]; comment:每日时间段"`
}

type QuickGenerateIn struct {
	ServiceID    int  `json:"service_id" binding:"required"`
	StartDate    Date `json:"start_date" binding:"required"`
	GenerateDays int  `json:"generate_days" binding:"required"`
}

type GetScheduleIn struct {
	ServiceID int  `json:"service_id" binding:"required"`
	Date      Date `json:"date" binding:"required"`
}
