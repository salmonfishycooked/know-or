package routes

import (
	"github.com/gin-gonic/gin"
	"go_web_app/logger"
	"net/http"
)

// Setup 用来初始化路由
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
