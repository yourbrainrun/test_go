package one

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	client2 "github.com/yourbrainrun/test_go/test_lua/client"
)

func ExecLuaOne() {
	ctx := context.Background()
	script := `return redis.call("keys",KEYS[1])`

	rdb := client2.GetRedisClient()

	luaCli := redis.NewScript(script)
	i, err := luaCli.Run(ctx, rdb, []string{"*"}).Result()
	if err != nil {
		return
	}

	fmt.Println(i)
}
