package two

import (
	"context"
	"fmt"
	client2 "github.com/yourbrainrun/test_go/test_lua/client"
)

func ExecTwo() {
	ctx := context.Background()
	script := `return redis.call("keys",KEYS[1])`

	rdb := client2.GetRedisClient()

	result, err := rdb.ScriptLoad(ctx, script).Result()
	if err != nil {
		return
	}

	ret, err := rdb.EvalSha(ctx, result, []string{"*"}).Result()
	if err != nil {
		return
	}
	fmt.Println(ret)
}
