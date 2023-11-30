package user

import (
	"encoding/base64"
	"encoding/json"
	"server/configs"
	"server/src/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 刷新 token
func RefreshToken(ctx *gin.Context) {
	type Params struct {
		Token string `binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams(ctx)
		return
	}

	info, err := service.Jwt.Verify(params.Token)
	if err != nil && err.Error() != "Token is expired" {
		service.State.ErrorUnauthorized(ctx, err.Error())
		return
	}

	startIndex := strings.Index(params.Token, ".") + 1
	endIndex := startIndex + strings.Index(params.Token[startIndex:], ".")
	tokenBody := params.Token[startIndex:endIndex]
	decodedData, _ := base64.StdEncoding.DecodeString(tokenBody)

	var data map[string]interface{}
	json.Unmarshal([]byte(decodedData), &data)

	// 超过__时间未登录，拒绝更新 token，通知退出
	poor := time.Now().Unix() - int64(data["exp"].(float64))
	if poor > configs.Config.TokenExceedRefreshTime {
		service.State.ErrorUnauthorized(ctx, "Leave too long, please log in again")
		return
	}

	token := service.Jwt.Publish(info)
	service.State.SuccessData(ctx, token)
}
