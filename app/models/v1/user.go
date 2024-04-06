package models_v1

type Status string
type Gender int
type Language string

const (
	Uncomplete Status = "uncomplete" // 未完善信息
	Complete   Status = "complete"   // 已完善信息

	Unknown Gender = 0 // 未知
	Male    Gender = 1 // 男
	Female  Gender = 2 // 女

	ZH_CN Language = "zh_CN" // 简体中文
	EN_US Language = "en_US" // 英文
	ZH_TW Language = "zh_TW" // 繁体中文

)

type User struct {
	BaseModel
	Nickname   string   `gorm:"type:varchar(20);"`
	Desciption string   `gorm:"default: null"`
	RoleID     int      `gorm:"default: null"`
	ServiceID  int      `gorm:"default: null"`
	OpenID     string   `gorm:"type:varchar(50);default: null"`
	Phone      string   `gorm:"type:varchar(50);default: null"`
	AvatarUrl  string   `gorm:"type:varchar(255);default: null"`
	Status     Status   `gorm:"type:varchar(20);default: uncomplete"`
	Province   string   `gorm:"type:varchar(50);default: null"`
	City       string   `gorm:"type:varchar(50);default: null"`
	Country    string   `gorm:"type:varchar(50);default: null"`
	Gender     Gender   `gorm:"type:int;default: 0"`
	Language   Language `gorm:"type:varchar(50);default: zh_CN"`
}

type UserSchema struct {
	ID        int      `json:"id"`
	Nickname  string   `json:"nickName,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	AvatarUrl string   `json:"avatarUrl,omitempty"`
	Status    Status   `json:"status,omitempty"`
	Province  string   `json:"province,omitempty"`
	City      string   `json:"city,omitempty"`
	Country   string   `json:"country,omitempty"`
	Gender    Gender   `json:"gender,omitempty"`
	Language  Language `json:"language,omitempty"`
}

// func (c BaseSchema) MarshalJSON() ([]byte, error) {
// 	var buf bytes.Buffer
// 	if len(string(c)) == 0 {
// 		buf.WriteString(`null`)
// 	} else {
// 		buf.WriteString(`"` + string(c) + `"`) // add double quation mark as json format required
// 	}
// 	return buf.Bytes(), nil
// }

// func (c *BaseSchema) UnmarshalJSON(in []byte) error {
// 	str := string(in)
// 	if str == `null` {
// 		*c = ""
// 		return nil
// 	}
// 	res := MyType(str)
// 	if len(res) >= 2 {
// 		res = res[1 : len(res)-1] // remove the wrapped qutation
// 	}
// 	*c = res
// 	return nil
// }
