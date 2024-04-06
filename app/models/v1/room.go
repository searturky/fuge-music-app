package models_v1

type Room struct {
	BaseModel
	Name       string `gorm:"type:varchar(50);"`
	StoreID    int    `gorm:"type:int; default:null"`
	ImageUrls  string `gorm:"type:text; default:null; comment:'图片链接, 逗号分隔'"`
	Desciption string `gorm:"type:text; default:null"`
}

type RoomSchema struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	StoreID    int    `json:"storeId,omitempty"`
	ImageUrls  string `json:"imageUrls,omitempty"`
	Desciption string `json:"desciption,omitempty"`
}
