package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Get 获取缓存字符串
func Get(ctx context.Context, key string) (string, error) {
	result, err := Cache.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return result, nil
}
