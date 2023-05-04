package mysql

import "go_web_app/model"

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

func GetPostList(page, pageSize int64) (data []*model.Post, err error) {
	sqlStr := `SELECT 
    post_id, title, content, author_id, community_id, status, create_time
    FROM post
    LIMIT ?, ?`
	data = make([]*model.Post, 0, 2)
	err = db.Select(&data, sqlStr, (page-1)*pageSize, pageSize)
	return
}
