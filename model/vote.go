package model

// VoteRes 投票接口返回数据
type VoteRes struct {
	Supports  int64 `json:"supports"`
	Direction int8  `json:"direction"`
}
