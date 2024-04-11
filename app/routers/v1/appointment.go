package router_v1

import (
	"net/http"
	"strconv"

	models "fuge/app/models/v1"
	services "fuge/app/service/v1"

	"github.com/gin-gonic/gin"
)

func AppointmentRouter(routerGroup *gin.RouterGroup) {
	appointmentGroup := routerGroup.Group("/appointment")
	getAppointmentByUserId(appointmentGroup) // 获取预定时间
	createAppointment(appointmentGroup)      // 创建预定时间
}

// @Summary 获取预定时间
// @Description 通过用户ID获取预定时间
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /appointment/{id} [get]
// @Param id path string true "用户ID"
func getAppointmentByUserId(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/:id",
		func(c *gin.Context) {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			c.JSON(200, gin.H{
				"message": id,
			})
		},
	)
}

// @Summary 创建预定时间
// @Description 创建预定时间
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /appointment [post]
// @Param param body models.CreateBookingIn true "预定日期"
func createAppointment(routerGroup *gin.RouterGroup) {
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
