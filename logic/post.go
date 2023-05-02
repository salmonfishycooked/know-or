package logic

import (
	"go.uber.org/zap"
	"go_web_app/dao/mysql"
	"go_web_app/model"
	"go_web_app/pkg/snowflake"
)

// CreatePost 用来创建一条帖子
func CreatePost(p *model.Post) (err error) {
	// 查询发帖的社区id是否存在
	if _, err = mysql.GetCommunityDetailByID(p.CommunityID); err != nil {
		return err
	}
	// 生成帖子id
	p.ID = snowflake.GenID()
	// 插入数据库
	return mysql.CreatePost(p)
}

// GetPostByID 返回对应id的帖子详情
func GetPostByID(pid int64) (*model.ApiPostDetail, error) {
	// 查询帖子详情
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return nil, err
	}

	// 查询作者
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return nil, err
	}

	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed",
			zap.Int64("author_id", post.CommunityID),
			zap.Error(err))
		return nil, err
	}

	// 返回数据
	data := &model.ApiPostDetail{
		AuthorName:      user.Username,
		CommunityDetail: community,
		Post:            post,
	}
	return data, err
}
