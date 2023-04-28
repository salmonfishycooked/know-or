package routes

import (
	"github.com/gin-gonic/gin"
	"go_web_app/controller"
	"go_web_app/logger"
)

// Setup 用来初始化路由
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 用户业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	return r
}
