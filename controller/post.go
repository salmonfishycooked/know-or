package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"know_or/logic"
	"know_or/model"
	"know_or/pkg/e"
	"know_or/pkg/utils"
	"strconv"
)

const (
	defaultPage     = 1
	defaultPageSize = 8
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
		return
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
	var data *model.ApiPostDetail
	err = utils.SetCurrentUserWithCookie(c)
	if err != nil {
		data, err = logic.GetPostByID(pid)
	} else {
		uid, _ := GetCurrentUser(c)
		data, err = logic.GetPostByIDWithUid(pid, uid)
	}

	if err != nil {
		zap.L().Error("logic.GetPostByID failed", zap.Error(err))
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回成功响应
	e.ResponseSuccess(c, data)
}

// GetPostListHandler 处理获取帖子列表的请求 新版
// 根据前端传来的参数动态获取帖子列表
// 按创建时间或者分数排序
func GetPostListHandler(c *gin.Context) {
	// 参数校验
	p := &model.ParamPostList{
		Page:     defaultPage,
		PageSize: defaultPageSize,
		Order:    model.OrderTime,
	}
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("GetPostListHandler with invalid param", zap.Error(err))
		e.ResponseError(c, e.CodeInvalidParam)
		return
	}

	// 获取帖子数据
	var data []*model.ApiPostDetail
	err := utils.SetCurrentUserWithCookie(c)
	if err != nil {
		data, err = logic.GetPostListNew(p)
	} else {
		uid, _ := GetCurrentUser(c)
		data, err = logic.GetPostListWithUid(p, uid)
	}

	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回成功响应
	e.ResponseSuccess(c, data)
}

// GetCommunityPostListHandler 查询指定社区的帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	// 参数校验
//	p := &model.ParamCommunityPostList{
//		ParamPostList: &model.ParamPostList{
//			Page:     defaultPage,
//			PageSize: defaultPageSize,
//			Order:    model.OrderTime,
//		},
//	}
//	if err := c.ShouldBind(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid param", zap.Error(err))
//		e.ResponseError(c, e.CodeInvalidParam)
//		return
//	}
//
//	// 获取帖子数据
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList() failed", zap.Error(err))
//		e.ResponseError(c, e.CodeServerBusy)
//		return
//	}
//
//	// 返回成功响应
//	e.ResponseSuccess(c, data)
//}
