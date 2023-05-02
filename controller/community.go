package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_web_app/logic"
	"go_web_app/pkg/e"
	"strconv"
)

// CommunityHandler 获取社区列表
func CommunityHandler(c *gin.Context) {
	// 将查询到所有的社区 (community_id,  community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		e.ResponseError(c, e.CodeServerBusy)
		return
	}
	e.ResponseSuccess(c, &data)
}

// CommunityDetailHandler 获取社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		e.ResponseError(c, e.CodeInvalidParam)
		return
	}

	// 将查询到所有的社区 (community_id,  community_name) 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		if err == e.ErrorInvalidID {
			e.ResponseError(c, e.CodeInvalidID)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
	}
	e.ResponseSuccess(c, &data)
}
