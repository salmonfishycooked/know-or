package redis

import "go_web_app/model"

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
