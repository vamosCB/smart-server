package cache

import (
	"smart-server/model/conf"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

// InitCache 初始化redis缓存
func InitCache() error {
	Cache = redis.NewClient(&redis.Options{
		Addr:        conf.RedisConf.Host,
		DialTimeout: time.Duration(conf.RedisConf.DialTimeout) * time.Second,
		MaxRetries:  conf.RedisConf.MaxRetries,
	})
	return nil
}
