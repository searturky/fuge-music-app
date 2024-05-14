package models_v1

type Room struct {
	BaseModel
	Name        string    `gorm:"type:varchar(50); comment: 房间名称"`
	StoreID     int       `gorm:"type:int; required; not null; index; comment: 店铺ID"`
	ImageUrls   []string  `gorm:"type:varchar[]; comment:图片链接, 逗号分隔"`
	Description string    `gorm:"type:text; default:null; comment: 房间描述"`
	Services    []Service `gorm:"many2many:room_services;"`
}

type RoomSchema struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	StoreID     int    `json:"storeId,omitempty"`
	ImageUrls   string `json:"imageUrls,omitempty"`
	Description string `json:"desciption,omitempty"`
}
