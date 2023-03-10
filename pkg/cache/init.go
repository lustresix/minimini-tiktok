package cache

import (
	"context"
	"fmt"
	"mini-min-tiktok/pkg/consts"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/spf13/viper"
)

var (
	RedisCache      *persist.RedisStore
	cacheMiddleware app.HandlerFunc
)

func Init() {
	// 连接设置
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: consts.TCP,
		Addr:    fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
	}))
	// 连接redis
	RedisCache = redisStore
	_, err := redisStore.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis 连接失败" + err.Error())
	}

	cacheMiddleware = cache.NewCacheByRequestURI(redisStore, 10*time.Minute)
}

func RedisMiddleware() app.HandlerFunc {
	return cacheMiddleware
}
