package main

import (
	"server/configs"
	v1 "server/src/controller/v1"
	"server/src/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func server() *gin.Engine {

	app := gin.Default()
	app.Use(middleware.Core)

	// 代理应用
	app.Any("/permissions/*path", middleware.ProxyPermissions)

	// 自身应用
	base := app.Group(configs.Config.Prefix)
	base.Use(middleware.Recover)
	base.Use(middleware.Logs)
	base.Use(middleware.BodyDispose)
	base.Use(middleware.Timeout)

	v1.Route(base.Group("/v1/api"))

	return app
}

func main() {

	port := ":" + strconv.Itoa(configs.Config.Port)
	server().Run(port)

}
