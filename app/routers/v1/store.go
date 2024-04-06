package router_v1

// import (
// 	"net/http"
// 	"strconv"

// 	models "fuge/app/models/v1"
// 	services "fuge/app/service/v1"

// 	"github.com/gin-gonic/gin"
// )

// @Summary 获取门店列表
// @Description 通过ID获取门店列表
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /booking/{id} [get]
// @Param id path string true "用户ID"
// func getAppointmentByUserId(routerGroup *gin.RouterGroup) {
// 	routerGroup.GET("/store/:id",
// 		func(c *gin.Context) {
// 			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 			c.JSON(200, gin.H{
// 				"message": id,
// 			})
// 		},
// 	)
// }
