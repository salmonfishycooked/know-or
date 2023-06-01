package logic

import (
	"go.uber.org/zap"
	"know_or/dao/redis"
	"know_or/model"
	"strconv"
)

// 投票功能

// VoteForPost 给某条帖子投票
func VoteForPost(userID int64, p *model.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
