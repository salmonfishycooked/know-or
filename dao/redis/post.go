package redis

import (
	"github.com/go-redis/redis"
	"go_web_app/model"
	"strconv"
	"time"
)

func getIDsFromKey(key string, page, pageSize int64) ([]string, error) {
	start := (page - 1) * pageSize
	end := start + pageSize - 1
	return client.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder 从缓存中获取相对应的帖子id
func GetPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	// 根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	return getIDsFromKey(key, p.Page, p.PageSize)
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

// GetCommunityPostIDsInOrder 按社区查询ids
func GetCommunityPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	cKey := getRedisKey(KeyCommunitySetPrefix + strconv.Itoa(int(p.CommunityID)))
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	// 若不存在这个 key
	if client.Exists(key).Val() < 1 {
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}

	return getIDsFromKey(key, p.Page, p.PageSize)
}
