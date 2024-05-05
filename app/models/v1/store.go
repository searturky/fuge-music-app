package models_v1

import (
// "gorm.io/gorm"
)

type Store struct {
	BaseModel
	Name       string `gorm:"type:varchar(50);"`
	Desciption string `gorm:"type:text; default:null"`
	Location   string `gorm:"type:varchar(50);"`
	Users      []User `gorm:"many2many:store_users;"`
}
