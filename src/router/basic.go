package router

import (
	"server/src/controller/user"

	"github.com/gin-gonic/gin"
)

func Basic(r *gin.RouterGroup) {
	r.POST("/login", user.SignIn)
	r.POST("/token/refresh", user.RefreshToken)
}
