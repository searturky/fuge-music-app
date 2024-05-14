package models_v1

type ServiceCategory struct {
	BaseModel
	StoreID  int    `gorm:"type:int; default:null; comment:门店ID"`
	Name     string `gorm:"type:varchar(50); comment:服务分类名称"`
	ParentID int    `gorm:"type:int; default:null; comment:父分类ID"`
}
