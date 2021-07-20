package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1 := context.WithValue(ctx, "value", "123123")

	go routineValue(ctx1)

	ctx2, function := context.WithCancel(ctx)
	go routineDone(ctx2)
	time.Sleep(time.Second * 3)
	function()

	time.Sleep(time.Second * 30)
}

func routineOne() {

}

func routineValue(ctx context.Context) {
	fmt.Println(ctx.Value("value"))
}

func routineDone(ctx context.Context) {
	ctx2 := context.WithValue(ctx, "level2", 3333)
	go routineDoneLevel2(ctx2)
	for true {
		select {
		case <-(ctx).Done():
			fmt.Println("routine Done over")
			return
		default:
			fmt.Println("default")
		}
		fmt.Println("running...")
		time.Sleep(time.Second * 1)
	}
}

func routineDoneLevel2(ctx context.Context) {
	for true {
		select {
		case <-(ctx).Done():
			fmt.Println("routine Done over routineDoneLevel2")
			return
		default:
			fmt.Println("routineDoneLevel2 default")
			fmt.Println(ctx.Value("level2"))
		}
		time.Sleep(time.Second * 1)
	}
}
