package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

type stateType struct {
	Code    int
	Data    any
	Message string
	RunTime string
}

// 贯穿整个中间件的数据
var State = stateType{}

func BodyDispose(ctx *gin.Context) {
	startTime := time.Now()
	State.Code = 400
	State.Message = "unknown error"
	State.RunTime = startTime.String()

	ctx.Next()

	// 如果已经返回了结果，不对数据进行包装
	if ctx.Writer.Written() {
		return
	}

	if State.Data != "" && State.Data != nil {
		State.Code = 200
		State.Message = "success"
	}

	State.RunTime = time.Since(startTime).String()
	ctx.JSON(200, gin.H{
		"code":    State.Code,
		"data":    State.Data,
		"message": State.Message,
		"runTime": State.RunTime,
	})

}
