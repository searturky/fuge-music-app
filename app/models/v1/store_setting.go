package models_v1

type StoreSetting struct {
	BaseModel
	StoreID int `gorm:"required; not null; index"`
	// BookingPeriod int    `gorm:"type:int;"`
	// BookingUnit   string `gorm:"type:varchar(10);"`
	// StartTime   string `gorm:"type:varchar(10);"`
	// EndTime     string `gorm:"type:varchar(10);"`
	// OpenWeekday []int  `gorm:"type:integer[];"` // 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
}
