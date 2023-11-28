package basic

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.POST("/login", SignIn)
}
