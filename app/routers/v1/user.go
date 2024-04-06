package router_v1

import (
	models "fuge/app/models/v1"
	services "fuge/app/service/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func loginIn(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/loginIn",
		func(c *gin.Context) {
			// 获取客户端是否携带cookie
			_, err := c.Cookie("key_cookie")
			if err != nil {
				c.SetCookie("key_cookie", "value_cookie", // 参数1、2： key & value
					10,          // 参数3： 生存时间（秒）
					"/",         // 参数4： 所在目录
					"localhost", // 参数5： 域名
					false,       // 参数6： 安全相关 - 是否智能通过https访问
					true,        // 参数7： 安全相关 - 是否允许别人通过js获取自己的cookie
				)
				c.String(200, "login success")
				return
			}
			c.String(200, "already login")
		},
	)
}

// 尝试访问，添加身份认证中间件，如果已经登陆就可以执行
func sayHello(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/sayHello", func(c *gin.Context) {
		c.String(200, "hello")
	})
}

// @Summary 登陆微信
// @Description 登陆微信
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /login-wechat [post]
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
