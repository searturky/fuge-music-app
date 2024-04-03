package models_v1

// "gorm.io/gorm"

type User struct {
	BaseModel
	Nickname   string `gorm:"type:varchar(20);"`
	Desciption string `gorm:"default: null"`
	RoleID     uint   `gorm:"default: null"`
	ServiceID  uint   `gorm:"default: null"`
	OpenID     string `gorm:"type:varchar(50);"`
	Phone      string `gorm:"type:varchar(50);default: null"`
	AvatarUrl  string `gorm:"type:varchar(255);default: null"`
}
