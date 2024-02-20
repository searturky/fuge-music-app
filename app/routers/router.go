package routers

import (
	v1 "fuge/app/routers/v1"
	"github.com/gin-gonic/gin"
)

type Register func(*gin.RouterGroup)

// InitRouter initialize routing information
func InitRouter(e *gin.Engine) {
	apiGroup := e.Group("/api")
	rs := [...]Register{
		v1.Router,
	}

	for _, r := range rs {
		r(apiGroup)
	}
}
