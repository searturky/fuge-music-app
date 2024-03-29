package models_v1

import (
	// "gorm.io/gorm"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Service struct {
	BaseModel
	Name       string `gorm:"type:varchar(50);"`
	Desciption string
	TimePeriod uint16 `gorm:"default:1; check:time_period > 0 and time_period < 1440 and 3600 % time_period = 0"`
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
