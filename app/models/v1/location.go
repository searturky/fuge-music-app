package models_v1

type Location struct {
	BaseModel
	Name       string
	Desciption string
	StoreID    uint
	Store      Store `gorm:"foreignKey:StoreID"`
}
