package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"know_or/logic"
	"know_or/model"
	"know_or/pkg/e"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := &model.ParamVoteData{}
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e.ResponseError(c, e.CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		e.ResponseErrorWithMsg(c, e.CodeInvalidParam, errData)
		return
	}

	// 获取当前发起请求的用户的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		e.ResponseError(c, e.CodeNeedLogin)
		return
	}
	// 投票
	data, err := logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		if err == e.ErrorVoteRepeat {
			e.ResponseError(c, e.CodeVoteRepeat)
			return
		} else if err == e.ErrorVoteTimeExpire {
			e.ResponseError(c, e.CodeVoteTimeExpire)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回成功响应
	e.ResponseSuccess(c, &data)
}
