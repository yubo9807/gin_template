package v1

import (
	"server/src/controller/v1/test"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.GET("/test", test.Gain)
}
