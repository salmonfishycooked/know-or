package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_web_app/middleware"
	"go_web_app/pkg/e"
)

var ErrorNeedLogin = errors.New(e.CodeNeedLogin.Msg())

// GetCurrentUser 用来获取当前请求的用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middleware.CtxUserIDKey)
	if !ok {
		err = ErrorNeedLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorNeedLogin
		return
	}
	return
}
