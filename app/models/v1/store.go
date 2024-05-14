package models_v1

// "gorm.io/gorm"

type Store struct {
	BaseModel
	Name       string   `gorm:"type:varchar(50); comment: 店铺名称"`
	Desciption string   `gorm:"type:text; default:null; comment: 店铺描述"`
	Location   string   `gorm:"type:varchar(50); comment: 店铺位置"`
	Users      []User   `gorm:"many2many:store_users;"`
	ImageUrls  []string `gorm:"type:varchar[]; comment:'图片链接, 逗号分隔'"`
}
