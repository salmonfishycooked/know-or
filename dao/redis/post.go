package redis

import (
	"github.com/go-redis/redis"
	"go_web_app/model"
)

// GetPostIDsInOrder 从缓存中获取相对应的帖子id
func GetPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	// 按分数从大到小查询
	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize - 1
	return client.ZRevRange(key, start, end).Result()
}

// GetPostSupportData 根据ids查询每篇帖子的投赞成票的数据
func GetPostSupportData(ids []string) ([]int64, error) {
	// 使用pipeline一次查完，减少网络RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	data := make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}
