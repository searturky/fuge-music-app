package router_v1

import (
	models "fuge/app/models/v1"
	services "fuge/app/service/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/user")
	loginWechat(userGroup) // 登陆微信
}

// @Summary 登陆微信
// @Description 登陆微信
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/login-wechat [post]
// @Param param body models.LoginWeChatIn true "登陆参数"
func loginWechat(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/login-wechat", func(c *gin.Context) {
		lwi := &models.LoginWeChatIn{}
		if err := c.ShouldBindJSON(lwi); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ret, err := services.WechatService.LoginWechat(lwi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(200, ret)
	})
}
