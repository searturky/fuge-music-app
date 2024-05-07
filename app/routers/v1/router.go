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
	WorkRouter(v1Group)        // 工作
	UserRouter(v1Group)        // 用户
	RoomRouter(v1Group)        // 房间
	AppointmentRouter(v1Group) // 预约
	ScheduleRouter(v1Group)    // 排班
}
