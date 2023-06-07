package utils

import (
	"github.com/gin-gonic/gin"
	"know_or/pkg/e"
	"know_or/pkg/jwt"
	"know_or/settings"
)

const (
	CtxUserIDKey = "userID"
	CtxUserToken = "userToken"
)

// SetCurrentUserWithCookie 从 Cookie 中读数据并存储到 context 里
func SetCurrentUserWithCookie(c *gin.Context) error {
	// 从 cookie 中获取 token
	token, err := c.Cookie(settings.COOKIE_TOKEN_FIELD)
	if err != nil {
		return e.ErrorNeedLogin
	}

	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := jwt.ParseToken(token)
	if err != nil {
		return e.ErrorInvalidToken
	}

	// 将当前请求的username信息保存到请求的上下文c上
	c.Set(CtxUserIDKey, mc.UserID)
	// 将当前请求的user token信息保存到请求的上下文c上
	c.Set(CtxUserToken, token)
	return nil
}
