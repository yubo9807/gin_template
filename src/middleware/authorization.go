package middleware

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

const KEY = "user_info"

func Authorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		service.State.ErrorUnauthorized(ctx)
		ctx.Abort()
		return
	}

	info, err := service.Jwt.Verify(auth)
	if err != nil {
		service.State.ErrorCustom(ctx, err.Error())
		ctx.Abort()
		return
	}
	ctx.Set(KEY, info)
}

// 获取 token 储存信息
func GetTokenInfo(ctx *gin.Context) map[string]interface{} {
	info, _ := ctx.Get(KEY)
	return info.(map[string]interface{})
}
