package routes

import (
	"github.com/gin-gonic/gin"
	"know_or/controller"
	"know_or/logger"
	"know_or/middleware"
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

	r.Use(middleware.Cors())

	v1 := r.Group("/api/v1")

	// 令牌桶限流中间件
	v1.Use(middleware.RateLimitMiddleware(BUCKET_RATE, BUCKET_CAP))

	// 用户业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	post := v1.Group("/post")
	{
		post.GET("/:id", controller.GetPostDetailHandler)
		// 根据时间或分数获取帖子列表
		post.GET("", controller.GetPostListHandler)

		// 对以下接口应用中间件
		post.Use(middleware.JWTAuthMiddleware())
		post.Use(middleware.IsRepeatLoginMiddleware())

		// 投票
		post.POST("/vote", controller.PostVoteHandler)
		// 新建帖子
		post.POST("", controller.CreatePostHandler)
	}

	community := v1.Group("/community")
	{
		community.GET("", controller.CommunityHandler)
		community.GET("/:id", controller.CommunityDetailHandler)
	}

	return r
}
