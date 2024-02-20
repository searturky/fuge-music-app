package router_v1

import (
	"net/http"

	models "fuge/app/models/v1"
	services "fuge/app/service/v1"

	"github.com/gin-gonic/gin"
)

// @Summary 创建服务
// @Description 创建服务
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /serivce [post]
// @Param param body models.CreateServiceIn true "创建服务参数"
func createService(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/serivce", func(c *gin.Context) {
		var csi = &models.CreateServiceIn{}
		if err := c.ShouldBindJSON(csi); err != nil {
			a := err.Error()
			print(a)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		services.ServiceService.CreateService(csi)
		c.Status(http.StatusOK)
	},
	)
}
