package models_v1

type Schedule struct {
	BaseModel
	StoreID           int    `gorm:"type:int; default:null"`
	RoomID            int    `gorm:"type:int; default:null"`
	ServiceCategoryID int    `gorm:"type:int; default:null"`
	ServiceID         int    `gorm:"type:int; default:null"`
	UserID            int    `gorm:"type:int; default:null"`
	Date              string `gorm:"type:varchar(30);"`
	StartTime         string `gorm:"type:varchar(30);"`
	EndTime           string `gorm:"type:varchar(30);"`
	TimePeriod        uint   `gorm:"required; not null;"`
	TimeSlot          string `gorm:"type:varchar(30);"`
}
