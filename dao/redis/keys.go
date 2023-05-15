package redis

// redis key 使用命名空间的方式，方便查询和拆分

const (
	keyPrefix = "bluebell:"

	KeyPostTimeZSet  = "post:time"  // zset; 帖子及发帖时间
	KeyPostScoreZSet = "post:score" // zset; 帖子及投票的分数

	KeyPostVotedZSetPrefix = "post:voted:" // zset; 记录用户及投票的类型
	KeyCommunitySetPrefix  = "community:"  // set; 保存每个分区下的帖子id
	KeyUserTokenPrefix     = "token:"      // string; 保存登录用户对应的token
)

// getRedisKey 给 redis key 加上前缀
func getRedisKey(key string) string {
	return keyPrefix + key
}
