package models_v1

import (
	"encoding/json"
	"time"
)

type Date time.Time

var _ json.Unmarshaler = &Date{}

const dateFormat = "2006-01-02"

func (mt *Date) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(dateFormat, s, time.Local)
	if err != nil {
		return err
	}
	*mt = Date(t)
	return nil
}

func (mt *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*mt).Format(dateFormat))
}

func (mt *Date) String() string {
	return time.Time(*mt).Format(dateFormat)
}

func (mt *Date) Time() time.Time {
	return time.Time(*mt)
}

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime; default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime; default:now()"`
	DeletedAt time.Time `gorm:"default: null"`
	Disabled  bool      `gorm:"default:false"`
}

var Models = []interface{}{
	&Booking{},
	&Location{},
	&Store{},
	&User{},
	&Service{},
	&Role{},
}
