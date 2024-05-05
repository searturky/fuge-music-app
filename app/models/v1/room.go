package models_v1

type Room struct {
	BaseModel
	Name        string `gorm:"type:varchar(50);"`
	StoreID     int    `gorm:"type:int; required; not null; index"`
	ImageUrls   string `gorm:"type:text; default:null; comment:'图片链接, 逗号分隔'"`
	Description string `gorm:"type:text; default:null"`

	Services []Service `gorm:"many2many:room_services;"`
}

type RoomSchema struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	StoreID     int    `json:"storeId,omitempty"`
	ImageUrls   string `json:"imageUrls,omitempty"`
	Description string `json:"desciption,omitempty"`
}
