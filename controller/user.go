package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_web_app/logic"
	"go_web_app/model"
	"go_web_app/pkg/e"
	"go_web_app/pkg/jwt"
	"go_web_app/settings"
)

// SignUpHandler 用来处理注册请求
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := &model.ParamSignUp{}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断 err 是否为 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e.ResponseError(c, e.CodeInvalidParam)
			return
		}
		e.ResponseErrorWithMsg(c, e.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, e.ErrorUserExist) {
			e.ResponseError(c, e.CodeUserExist)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 返回响应
	e.ResponseSuccess(c, nil)
}

// LoginHandler 用来处理登录请求
func LoginHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := &model.ParamLogin{}
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("login with invalid param", zap.Error(err))
		// 判断 err 是否为 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			e.ResponseError(c, e.CodeInvalidParam)
			return
		}
		e.ResponseErrorWithMsg(c, e.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, e.ErrorUserNotExist) {
			e.ResponseError(c, e.CodeUserNotExist)
			return
		} else if errors.Is(err, e.ErrorInvalidPassword) {
			e.ResponseError(c, e.CodeInvalidPassword)
			return
		}
		e.ResponseError(c, e.CodeServerBusy)
		return
	}

	// 设置Cookie token字段
	c.SetCookie(settings.COOKIE_TOKEN_FIELD, user.Token, int(jwt.TokenExpireDuration), "/", "localhost", false, true)

	// 返回响应
	e.ResponseSuccess(c, gin.H{
		"user_id":  user.UserID,
		"username": user.Username,
	})
}
