package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_app/pkg/e"
)

const CtxUserIDKey = "userID"

// GetCurrentUser 用来获取当前请求的用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = e.ErrorNeedLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = e.ErrorNeedLogin
		return
	}
	return
}
