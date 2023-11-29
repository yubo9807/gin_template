package router

import (
	"server/src/controller/test"
	"server/src/middleware"

	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	r.Use(middleware.Authorization)

	r.GET("/test", test.Test)
}
