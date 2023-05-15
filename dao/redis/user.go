package redis

import (
	"fmt"
	"strconv"
	"time"
)

// GetUserToken 获取对应uid的token
func GetUserToken(uid int64) (string, error) {
	key := KeyUserTokenPrefix + strconv.Itoa(int(uid))
	result, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	fmt.Println(result)
	return result, nil
}

// SetUserToken 在缓存中存入对应uid的token
// 供重复登录中间件校验用
func SetUserToken(uid int64, token string, expiration time.Duration) error {
	key := KeyUserTokenPrefix + strconv.Itoa(int(uid))
	_, err := client.Set(key, token, expiration).Result()
	return err
}
