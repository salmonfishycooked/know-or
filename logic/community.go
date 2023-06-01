package logic

import (
	"know_or/dao/mysql"
	"know_or/model"
)

// GetCommunityList 返回 community 列表
func GetCommunityList() ([]*model.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 返回指定社区id的详情信息
func GetCommunityDetail(id int64) (*model.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
