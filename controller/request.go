package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_app/pkg/e"
	"strconv"
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

// getPageInfo 用来统一处理前端获取列表时候的 page, page_size
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = 8
	}
	return page, pageSize
}
