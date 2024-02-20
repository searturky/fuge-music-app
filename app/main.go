package main

import (
	"fuge/app/core"
	"net/http"

	"fuge/app/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var conf *core.ConfYaml

func init() {
	conf = core.InitConf()
	core.InitDB(conf)
}

func main() {
	gin.SetMode(conf.GetGinMode())
	addr := conf.GetAddr()
	engine := gin.Default()
	engine.Use(cors.Default())
	routers.InitRouter(engine)
	engine.GET("/swagger/v1/*any", ginSwagger.WrapHandler(
		swaggerFiles.NewHandler(),
		ginSwagger.InstanceName("v1"),
	))
	s := &http.Server{
		Addr:    addr,
		Handler: engine,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	defer core.CloseDB()
	defer s.Close()
	s.ListenAndServe()
}
