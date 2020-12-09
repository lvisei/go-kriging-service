package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/liuvigongzuoshi/go-kriging-service/internal/app/config"
)

// RateLimiterMiddleware 请求频率限制中间件
func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.C.RateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		// TODO：

		c.Next()
	}
}
