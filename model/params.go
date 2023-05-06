package model

const (
	OrderTime  = "time"  // 按时间最新排序
	OrderScore = "score" // 按分数高低排序
)

// 定义请求的参数结构体

// ParamSignUp 用户注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票请求参数
type ParamVoteData struct {
	// UserID 从请求中可以获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=-1 0 1"` // 赞成(1) 反对(-1) 取消投票(0)
}

// ParamPostList 获取帖子列表参数 (query string参数)
type ParamPostList struct {
	Page     int64  `json:"page" form:"page"`
	PageSize int64  `json:"page_size" form:"page_size"`
	Order    string `json:"order" form:"order"`
}
