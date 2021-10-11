package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yourbrainrun/test_go/test_global_redis_lock/drivers"
	"github.com/yourbrainrun/test_go/test_global_redis_lock/helpers"
	"time"
)

func init() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
	}
}

func main() {
	for i := 10; i > 0; i-- {
		go test()
	}
	time.Sleep(1 * time.Second)

	//list()
}

func list() {
	ctx := context.Background()
	rdb := drivers.GetRedisClient()
	key := "list:push"
	rdb.LPush(ctx,key,"124.0.9.23")
}

func test() {
	key := "eqweqweqweqweqwe1234"
	lock := &helpers.RedisLock{ReleaseDuration: 10 * time.Second}
	if lock.Lock(key) == false {
		fmt.Println("already")
		return
	}
	fmt.Println("locked")
	time.Sleep(10 * time.Second)
	defer lock.Unlock(key)
}
