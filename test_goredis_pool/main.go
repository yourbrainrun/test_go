package main

import (
	"context"
	"fmt"
	"github.com/yourbrainrun/test_go/test_goredis_pool/helpers"
)

func main() {
	for i := 0; i < 500; i++ {
		go testConn()
	}
}

func testConn() {
	key := "test:key:list"
	ctx := context.Background()
	redisPool := helpers.GetRedisClient()
	if str, err := redisPool.LPop(ctx, key).Result(); err == nil {
		fmt.Println(str)
	} else {
		fmt.Println(err.Error())
	}
}
