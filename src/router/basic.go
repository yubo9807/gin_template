package router

import (
	"server/src/controller/login"

	"github.com/gin-gonic/gin"
)

func Basic(r *gin.RouterGroup) {
	r.POST("/login", login.SignIn)
}
