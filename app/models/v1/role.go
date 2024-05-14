package models_v1

// "gorm.io/gorm"

type RoleEnum string

const (
	UserRole   RoleEnum = "user"
	AdminRole  RoleEnum = "admin"
	ServerRole RoleEnum = "server"
)

type Role struct {
	BaseModel
	Name       RoleEnum `gorm:"type:varchar(50); comment: 角色名称"`
	Desciption string   `gorm:"type:text; default:null; comment: 角色描述"`
}
