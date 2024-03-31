package models_v1

import (
// "gorm.io/gorm"
)

type User struct {
	BaseModel
	Nickname   string `gorm:"type:varchar(20);"`
	Desciption string
	RoleID     uint
	ServiceID  uint
	OpenID     string `gorm:"type:varchar(50);"`
	Phone      string `gorm:"type:varchar(50);"`
}
