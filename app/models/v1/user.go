package models_v1

import (
// "gorm.io/gorm"
)

type User struct {
	BaseModel
	Name       string
	Desciption string
	RoleID     uint
	ServiceID  uint
}
