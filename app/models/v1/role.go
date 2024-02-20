package models_v1

import (
// "gorm.io/gorm"
)

type Role struct {
	BaseModel
	Name       string
	Desciption string
}
