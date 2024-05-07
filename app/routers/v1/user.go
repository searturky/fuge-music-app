package router_v1

import (
	"fuge/app/middleware"
	models "fuge/app/models/v1"
	services "fuge/app/service/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	userGroup := routerGroup.Group("/user")
	loginWechat(userGroup)    // 登陆微信
	getPhoneNumber(userGroup) // 获取用户手机号
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
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		ret, err := services.WechatService.LoginWechat(lwi)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(200, ret)
	})
}

// @Summary 获取用户手机号
// @Description 获取用户手机号
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/phone-number [get]
// @Param code query string true "微信code"
// @Param Authorization header string true "Authorization"
func getPhoneNumber(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/phone-number", middleware.AuthUnCompleteMiddleWare(), func(c *gin.Context) {
		code := c.Query("code")
		data, exist := c.Get("user")
		if !exist {
			c.JSON(403, gin.H{})
			c.Abort()
			return
		}
		user := data.(*models.User)
		err := services.WechatService.GetPhoneNumber(code, user)
		if err != nil {
			c.JSON(500, gin.H{})
			return
		}
		c.JSON(200, gin.H{})
	})
}
