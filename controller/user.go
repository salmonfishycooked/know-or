package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go_web_app/logic"
	"go_web_app/model"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
