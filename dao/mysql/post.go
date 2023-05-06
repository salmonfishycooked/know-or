package mysql

import (
	"github.com/jmoiron/sqlx"
	"go_web_app/model"
	"strings"
)

// CreatePost 向数据库插入一条帖子
func CreatePost(p *model.Post) (err error) {
	sqlStr := `INSERT INTO
			post(post_id, title, content, author_id, community_id)
			VALUES(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostByID 向数据库查询对应id的帖子
func GetPostByID(pid int64) (post *model.Post, err error) {
	post = &model.Post{}
	sqlStr := `SELECT
			post_id, title, content, author_id, community_id, status, create_time
			FROM post
			WHERE post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表
func GetPostList(page, pageSize int64) (data []*model.Post, err error) {
	sqlStr := `SELECT 
    post_id, title, content, author_id, community_id, status, create_time
    FROM post
    ORDER BY create_time
    DESC
    LIMIT ?, ?`
	data = make([]*model.Post, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*pageSize, pageSize)
	return
}

// GetPostListByIDs 查询给定ids的帖子，顺序也按照给定的返回
func GetPostListByIDs(ids []string) (postList []*model.Post, err error) {
	sqlStr := `SELECT 
    post_id, title, content, author_id, community_id, status, create_time
	FROM post
	WHERE post_id in (?)
	ORDER BY FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
