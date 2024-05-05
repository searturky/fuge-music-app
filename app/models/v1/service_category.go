package models_v1

type ServiceCategory struct {
	BaseModel
	StoreID  int    `gorm:"type:int; default:null"`
	Name     string `gorm:"type:varchar(50);"`
	ParentID int    `gorm:"type:int; default:null"`
}
