package login

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

// 登录
func SignIn(ctx *gin.Context) {
	type Params struct {
		Username string `binding:"required"`
		Password string `binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	info := map[string]interface{}{
		"username": params.Username,
		"password": params.Password,
	}
	token := service.Jwt.Publish(info)
	service.State.SuccessData(ctx, token)
}
