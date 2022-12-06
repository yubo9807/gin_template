package main

import (
	v1 "server/src/controller/v1"
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func server() *gin.Engine {

	app := gin.Default()
	base := app.Group("/hicky")

	base.Use(middleware.Recover)
	base.Use(middleware.Core)
	base.Use(middleware.Logs)
	base.Use(middleware.BodyDispose)

	v1.Route(base.Group("/v1/api"))

	return app
}

func main() {

	server().Run(":20020")

}
