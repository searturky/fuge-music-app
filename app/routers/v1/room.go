package router_v1

import (
	"fuge/app/middleware"
	"net/http"
	"strconv"
	// models "fuge/app/models/v1"
	services "fuge/app/service/v1"

	"github.com/gin-gonic/gin"
)

func RoomRouter(routerGroup *gin.RouterGroup) {
	roomGroup := routerGroup.Group("/room", middleware.AuthMiddleWare())
	getRoomByStoreId(roomGroup) // 获取房间
}

// @Summary 获取门店的房间列表
// @Description 通过ID获取门店的房间列表
// @Tags v1, 门店相关
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /room/store/{storeId} [get]
// @Param storeId path int true "门店ID"
func getRoomByStoreId(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/store/:storeId",
		func(c *gin.Context) {
			store_id, err := strconv.Atoi(c.Param("storeId"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			services.RoomService.GetRoomByStoreId(store_id)
			c.JSON(200, gin.H{
				"message": store_id,
			})
		},
	)
}
