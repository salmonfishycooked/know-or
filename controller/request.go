package controller

import (
	"github.com/gin-gonic/gin"
	"know_or/pkg/e"
	"know_or/pkg/utils"
	"strconv"
)

// GetCurrentUser 用来获取当前请求的用户id
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(utils.CtxUserIDKey)
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

// GetCurrentUserToken 用来获取当前请求的用户token
func GetCurrentUserToken(c *gin.Context) (token string, err error) {
	data, ok := c.Get(utils.CtxUserIDKey)
	if !ok {
		err = e.ErrorNeedLogin
		return
	}
	token, ok = data.(string)
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
		page = defaultPage
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		pageSize = defaultPageSize
	}
	return page, pageSize
}
