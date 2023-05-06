package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go_web_app/settings"
)

var client *redis.Client

// Init 用来初始化与 Redis 的连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 使用的数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	_, err = client.Ping().Result()
	return
}

// Close 用来关闭与 Redis 的连接
func Close() {
	_ = client.Close()
}
