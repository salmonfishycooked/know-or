package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web_app/logic"
	"go_web_app/model"
	"go_web_app/pkg/e"
)

// CreatePostHandler 用来处理创建帖子请求
func CreatePostHandler(c *gin.Context) {
	// 参数校验
	p := &model.Post{}
	if err := c.ShouldBindJSON(p); err != nil {
		e.ResponseError(c, e.CodeInvalidParam)
		return
	}

	// 获取发帖用户id
	uid, err := GetCurrentUser(c)
	if err == e.ErrorNeedLogin {
		e.ResponseError(c, e.CodeNeedLogin)
		return
	}
	p.AuthorID = uid

	// 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
	}

	// 返回成功响应
	e.ResponseSuccess(c, nil)
}
