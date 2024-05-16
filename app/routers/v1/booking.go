package router_v1

import (
	"net/http"

	models "fuge/app/models/v1"
	services "fuge/app/service/v1"

	"github.com/gin-gonic/gin"
)

func BookingRouter(routerGroup *gin.RouterGroup) {
	appointmentGroup := routerGroup.Group("/booking")
	createBooking(appointmentGroup) // 创建预定
}

// @Summary 创建预定
// @Description 创建预定
// @Tags v1, 预定相关
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /booking [post]
// @Param param body models.CreateBookingIn true "预定日期"
func createBooking(routerGroup *gin.RouterGroup) {
	routerGroup.POST("", func(c *gin.Context) {
		var cbi = &models.CreateBookingIn{}
		if err := c.ShouldBindJSON(cbi); err != nil {
			a := err.Error()
			print(a)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		services.BookingService.CreateBooking(cbi)
		c.JSON(200, gin.H{})
	},
	)
}
