package redis

import (
	"github.com/go-redis/redis"
	"know_or/model"
	"know_or/pkg/e"
	"math"
	"strconv"
	"time"
)

// 本项目使用简化版的投票分数

// 投票的几种情况
// direction = 1 时，有两种情况
// 	1. 之前没有投过票，现在投赞成票    -> 更新分数和投票记录 差值的绝对值：1 +432
// 	2. 之前投反对票，现在投赞成票     -> 更新分数和投票记录 差值的绝对值：2 +432*2
// direction = 0 时，有两种情况
// 	1. 之前投过赞成票，现在要取消投票   -> 更新分数和投票记录 差值的绝对值：1 -432
// 	2. 之前投反对票，现在要取消投票    -> 更新分数和投票记录 差值的绝对值：1 +432
// direction = -1 时，有两种情况
//	1. 之前没有投过票，现在投反对票    -> 更新分数和投票记录 差值的绝对值：1 -432
//	2. 之前投赞成票，现在投反对票      -> 更新分数和投票记录 差值的绝对值：2 -432*2

// 投票的限制
// 每个帖子自发表之日起一个星期之内允许用户投票，超过时间则不允许投票
// 1. 到期之后将 redis 中保存的赞成票数以及反对票数存储到 mysql 表中
// 2. 到期之后删除 KeyPostVotedZSetPrefix

const (
	oneWeekInSeconds = 7 * 24 * 3600
	ScorePerVote     = 432 // 每一票的分值
)

// CreatePost 将新创建的帖子的创建时间保存到redis
func CreatePost(postID int64, communityID int64) error {
	pipeline := client.TxPipeline() // 使用 redis 事务
	// 帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数（默认分数为创建的时间）
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 把帖子id加到社区的set里
	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)

	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) (*model.VoteRes, error) {
	// 判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val() // 取帖子发布时间
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return nil, e.ErrorVoteTimeExpire
	}

	// 更新分数
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val() // 查询用户对当前帖子的投票记录

	// 判断是否重复投票
	if value == ov {
		return nil, e.ErrorVoteRepeat
	}

	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票状态差值

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*ScorePerVote, postID)
	// 记录用户的投票数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value, // 用户对该帖子的投票状态 -1反对 1赞成
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	// 查询帖子点赞数
	sup, err := GetPostSupportData([]string{postID})
	if err != nil {
		return nil, err
	}
	data := &model.VoteRes{
		Supports:  sup[0],
		Direction: int8(value),
	}
	return data, nil
}
