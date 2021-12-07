package client

import "github.com/go-redis/redis/v8"

var Client *redis.Client

func GetRedisClient() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Client
}
