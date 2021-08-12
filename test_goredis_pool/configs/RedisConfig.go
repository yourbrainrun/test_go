package configs

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"sync"
)

type RedisOption struct {
	Addr     string
	Password string
	Db       int
}

var redisConfigOnce sync.Once
var RedisConfig *redis.Options
func GetRedisOption() *redis.Options {
	redisConfigOnce.Do(func() {
		db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
		RedisConfig = &redis.Options{
			Addr:     os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       db,
		}
	})
	return RedisConfig
}

