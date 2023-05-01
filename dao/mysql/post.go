package mysql

import "go_web_app/model"

func CreatePost(p *model.Post) (err error) {
	sqlStr := `INSERT INTO
			post(post_id, title, content, author_id, community_id)
			VALUES(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}
