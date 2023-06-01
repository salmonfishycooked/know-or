package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"know_or/model"
	"know_or/pkg/e"
)

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := `SELECT community_id, community_name FROM community`
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (data *model.CommunityDetail, err error) {
	data = &model.CommunityDetail{}
	sqlStr := `SELECT
			community_id, community_name, introduction, create_time
			FROM community
			WHERE community_id=?`
	if err = db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = e.ErrorInvalidID
		}
	}
	return
}
