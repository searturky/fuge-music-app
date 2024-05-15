package router_v1

import (
	models "fuge/app/models/v1"
	services "fuge/app/service/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScheduleRouter(routerGroup *gin.RouterGroup) {
	scheduleGroup := routerGroup.Group("/schedule")
	quickGenerate(scheduleGroup)
	getScheduleByUserAndDate(scheduleGroup)
}

// @Summary 快速生成排班
// @Description 快速生成排班
// @Tags v1, 排班相关
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /schedule/quick-generate [post]
// @Param param body models.QuickGenerateIn true "快速生成排班参数"
func quickGenerate(routerGroup *gin.RouterGroup) {
	// routerGroup.POST("/quick-generate", middleware.AuthMiddleWare(), func(c *gin.Context) {
	routerGroup.POST("/quick-generate", func(c *gin.Context) {
		qgi := &models.QuickGenerateIn{}
		if err := c.ShouldBindJSON(qgi); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		services.ScheduleService.QuickGenerate(qgi)
		c.JSON(200, gin.H{})
	})
}

// @Summary 获取排班
// @Description 获取排班
// @Tags v1, 排班相关
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /schedule [get]
// @Param param query models.GetScheduleIn true "获取排班参数"
func getScheduleByUserAndDate(routerGroup *gin.RouterGroup) {
	// routerGroup.POST("/quick-generate", middleware.AuthMiddleWare(), func(c *gin.Context) {
	routerGroup.GET("", func(c *gin.Context) {
		gsi := &models.GetScheduleIn{}
		if err := c.ShouldBindQuery(gsi); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		services.ScheduleService.GetScheduleByUserAndDate(gsi)
		c.JSON(200, gin.H{})
	})
}
