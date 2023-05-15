package middleware

import (
	"github.com/gin-gonic/gin"
	"go_web_app/controller"
	"go_web_app/dao/redis"
	"go_web_app/pkg/e"
)

// IsRepeatLogin 限制同一账号只能在一台设备上登录
func IsRepeatLogin() func(*gin.Context) {
	return func(c *gin.Context) {
		// 拿取该请求的用户uid
		uid, err := controller.GetCurrentUser(c)
		if err != nil {
			e.ResponseError(c, e.CodeNeedLogin)
			c.Abort()
			return
		}

		// 去缓存里面取对应的用户 token
		rToken, err := redis.GetUserToken(uid)
		if err != nil {
			e.ResponseError(c, e.CodeNeedLogin)
			c.Abort()
			return
		}
		// 取当前请求 token
		token, err := controller.GetCurrentUserToken(c)
		if err != nil {
			e.ResponseError(c, e.CodeNeedLogin)
			c.Abort()
			return
		}

		// 对比 token，一致则通过，不一致则返回错误
		if rToken != token {
			e.ResponseError(c, e.CodeLoginRepeat)
			c.Abort()
			return
		}
		c.Next()
	}
}
