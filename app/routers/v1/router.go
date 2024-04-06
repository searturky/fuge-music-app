package router_v1

import (
	"github.com/gin-gonic/gin"

	_ "fuge/app/docs"
)

// @title fuge API
// @version 0.1
// @description fuge API
// @host localhost:8080
// @BasePath /api/v1
func Router(routerGroup *gin.RouterGroup) {
	v1Group := routerGroup.Group("/v1")
	loginIn(v1Group)                // 登录
	sayHello(v1Group)               // sayHello
	getAppointmentByUserId(v1Group) // 获取预定时间
	createAppointment(v1Group)      // 创建预定时间
	createService(v1Group)          // 创建服务
	loginWechat(v1Group)            // 微信登录
	getRoomByStoreId(v1Group)       // 获取房间
}
