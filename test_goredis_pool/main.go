package main

import (
	"carina-exporter/helpers"
	"context"
	"fmt"
	"sync"
)

// go run main.go
func main() {

	testExists()
	return
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go testConn(&wg)
	}
	wg.Wait()
}

func testConn(wg *sync.WaitGroup) {
	defer wg.Done()
	key := "test:key:list"
	ctx := context.Background()
	redisPool := helpers.GetRedisClient()
	if str, err := redisPool.LPop(ctx, key).Result(); err == nil {
		fmt.Println(str)

	} else {
		fmt.Println("err")
		fmt.Println(err.Error())
	}
}

func testExists() {
	key := "rtn://124.0.9.23"
	ctx := context.Background()
	redisPool := helpers.GetRedisClient()
	if str, err := redisPool.Exists(ctx, key).Result(); err == nil {
		fmt.Println(str)

	} else {
		fmt.Println(str, "err")
		fmt.Println(err.Error())
	}
}
