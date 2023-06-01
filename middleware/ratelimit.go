package middleware

import (
	"github.com/gin-gonic/gin"
	"know_or/pkg/e"
	"know_or/pkg/ratelimit"
	"time"
)

// RateLimitMiddleware 令牌桶限流中间件
func RateLimitMiddleware(rate time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(rate, cap)
	return func(c *gin.Context) {
		if !bucket.Allow() {
			e.ResponseError(c, e.CodeServerBusy)
			c.Abort()
			return
		}
		c.Next()
	}
}
