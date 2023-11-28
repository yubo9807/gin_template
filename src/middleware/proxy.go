package middleware

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 权限应用代理
func ProxyPermissions(ctx *gin.Context) {
	targetURL, err := url.Parse("http://localhost:20020/")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to parse target URL")
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	ctx.Request.URL.Scheme = targetURL.Scheme
	ctx.Request.URL.Host = targetURL.Host
	ctx.Request.Host = targetURL.Host
	ctx.Request.Header.Set("Open-Id", "1hendj97f")

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
