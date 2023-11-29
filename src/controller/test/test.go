package test

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context) {
	service.State.SuccessData(ctx, "success")
}
