package middleware

import (
	"server/src/config"

	"github.com/gin-gonic/gin"
)

func Core(ctx *gin.Context) {
	origin := "http://hpyyb.cn"
	if config.Env.DEVELOPMENT {
		origin = "*"
	}

	ctx.Header("Access-Control-Allow-Origin", origin)
}
