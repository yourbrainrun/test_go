package helpers

import (
	"carina-exporter/configs"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var redisPool *redis.Client
var getRedisPoolOnce sync.Once
var tempOption redis.Options

func GetRedisClient() *redis.Client {
	getRedisPoolOnce.Do(func() {
		redisConfig := configs.GetRedisOption()

		tempOption = redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,

			PoolSize:     15,
			MinIdleConns: 10,

			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolTimeout:  5 * time.Second,

			IdleCheckFrequency: 60 * time.Second,
			IdleTimeout:        2 * time.Minute,
			MaxConnAge:         0 * time.Second,

			MaxRetries:      0,
			MinRetryBackoff: 8 * time.Millisecond,
			MaxRetryBackoff: 512 * time.Millisecond,
		}

		redisPool = redis.NewClient(&tempOption)
	})

	return redisPool
}
