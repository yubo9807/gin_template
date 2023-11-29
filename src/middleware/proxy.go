package middleware

import (
	"net/http/httputil"
	"net/url"
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func formatResult(ctx *gin.Context, code int, msg string) {
	ctx.JSON(200, gin.H{"code": code, "message": msg})
	ctx.Abort()
}

// 代理前权限认证
func ProxyAuthorization(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		formatResult(ctx, 401, "unauthorized")
		return
	}
	_, err := service.Jwt.Verify(auth)
	if err != nil {
		formatResult(ctx, 401, err.Error())
		return
	}
}

// 权限应用代理
func ProxyPermissions(ctx *gin.Context) {
	targetURL, err := url.Parse("http://localhost:20020/")
	if err != nil {
		formatResult(ctx, 500, err.Error())
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	ctx.Request.URL.Scheme = targetURL.Scheme
	ctx.Request.URL.Host = targetURL.Host
	ctx.Request.Host = targetURL.Host
	ctx.Request.Header.Set("Open-Id", "1hendj97f")

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
