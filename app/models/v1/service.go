package models_v1

import (
	// "gorm.io/gorm"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	BaseModel
	StoreID     int     `gorm:"required; not null; index; comment:门店ID"`
	CategoryID  int     `gorm:"required; not null; index; comment:服务分类ID"`
	Name        string  `gorm:"type:varchar(50); comment:服务名称"`
	Description string  `gorm:"type:text; default:null; comment:服务描述"`
	Price       float64 `gorm:"type:decimal(10,2); default:0.00; comment:价格"` // 10位整数, 2位小数
	// DailyStartTime string  `gorm:"type:varchar(30); comment:每日服务开始时间"`                                                               // "09:00"
	// DailyEndTime   string  `gorm:"type:varchar(30); comment:每日服务结束时间"`                                                               // "21:00"
	TimePeriod int `gorm:"default:1; check:time_period > 0 and time_period < 1440 and 3600 % time_period = 0; comment:时间间隔"` // 1-1440分钟, 3600分钟必须能被整除

	Users []User `gorm:"many2many:user_services;"`
	Rooms []Room `gorm:"many2many:room_services;"`
	// WorkWeekday []int  `gorm:"type:integer[];"` // 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
}

type CreateServiceIn struct {
	Name        string `json:"name" binding:"required" type:"string" format:"string" description:"服务名称"`
	Description string `json:"description" binding:"required" type:"string" format:"string" description:"服务描述"`
	TimePeriod  uint16 `json:"timePeriod" binding:"required,serviceTimePeriod" type:"integer" format:"integer" validate:"min=1,max=3600;" description:"服务时长,单位分钟,最小1分钟,最大3600分钟,且3600分钟必须能被整除" example:"60"`
}

var serviceTimePeriod validator.Func = func(fl validator.FieldLevel) bool {
	period, ok := fl.Field().Interface().(uint16)
	if !ok || period < 1 || period > 3600 || 3600%period != 0 {
		return false
	}
	return true
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("serviceTimePeriod", serviceTimePeriod)
	}
}
