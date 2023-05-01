package logic

import (
	"go_web_app/dao/mysql"
	"go_web_app/model"
	"go_web_app/pkg/snowflake"
)

// CreatePost 用来创建一条帖子
func CreatePost(p *model.Post) (err error) {
	// 生成帖子id
	p.ID = snowflake.GenID()
	// 插入数据库
	return mysql.CreatePost(p)
}
