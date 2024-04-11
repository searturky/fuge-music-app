package router_v1

import (
	"net/http"

	models "fuge/app/models/v1"
	services "fuge/app/service/v1"

	"github.com/gin-gonic/gin"
)

func WorkRouter(routerGroup *gin.RouterGroup) {
	workkGroup := routerGroup.Group("/work")
	createWork(workkGroup) // 新建工作
}

// @Summary 新建工作
// @Description 新建工作
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /work [post]
// @Param param body models.CreateServiceIn true "创建工作内容参数"
func createWork(routerGroup *gin.RouterGroup) {
	routerGroup.POST("", func(c *gin.Context) {
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
