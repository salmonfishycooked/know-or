package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web_app/logic"
	"go_web_app/model"
	"go_web_app/pkg/e"
	"strconv"
)

// CreatePostHandler 处理创建帖子请求
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
		if err == e.ErrorInvalidID {
			e.ResponseError(c, e.CodeInvalidID)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
	}

	// 返回成功响应
	e.ResponseSuccess(c, nil)
}

// GetPostDetailHandler 处理获取对应id帖子详情请求
func GetPostDetailHandler(c *gin.Context) {
	// 获取并校验参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetail with invalid param", zap.Error(err))
		e.ResponseError(c, e.CodeInvalidParam)
		return
	}

	// 取出相对应id的帖子数据
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID failed", zap.Error(err))
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回成功响应
	e.ResponseSuccess(c, &data)
}

// GetPostListHandler 处理获取帖子列表的请求
func GetPostListHandler(c *gin.Context) {
	// 参数校验
	page, pageSize := getPageInfo(c)

	// 获取帖子数据
	data, err := logic.GetPostList(page, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回成功响应
	e.ResponseSuccess(c, data)
}
