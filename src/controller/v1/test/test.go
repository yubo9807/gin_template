package test

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func Gain(ctx *gin.Context) {
	service.State.SuccessData(ctx, "success")
}
