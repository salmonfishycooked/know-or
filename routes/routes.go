package routes

import (
	"github.com/gin-gonic/gin"
	"go_web_app/controller"
	"go_web_app/logger"
	"go_web_app/middleware"
	"time"
)

const (
	BUCKET_CAP  = 1000
	BUCKET_RATE = 10 * time.Millisecond
)

// Setup 用来初始化路由
func Setup() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.Use(middleware.RateLimitMiddleware(BUCKET_RATE, BUCKET_CAP)) // 令牌桶限流中间件
	// 用户业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware()) // 应用 JWT 认证中间件
	v1.Use(middleware.IsRepeatLogin())     // 单点登录检查
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		//v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts", controller.GetPostListHandler)

		// 投票
		v1.POST("/vote", controller.PostVoteHandler)
	}

	return r
}
