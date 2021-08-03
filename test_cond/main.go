package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	mut := sync.Mutex{}
	cond := sync.NewCond(&mut)

	go func() {
		cond.L.Lock()
		cond.Wait()
		fmt.Fprintln(os.Stdout, "say hello :", time.Now().Unix())
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		cond.Wait()
		fmt.Fprintln(os.Stdout, "say wait :", time.Now().Unix())
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		cond.Wait()
		fmt.Fprintln(os.Stdout, "say can you see:", time.Now().Unix())
		cond.L.Unlock()
	}()

	time.Sleep(1 * time.Second)
	cond.Signal()
	time.Sleep(2 * time.Second)
	cond.Broadcast()
	time.Sleep(3 * time.Second)

	// 测试出现问题 ，如果没有睡眠 第一个 cond.Signal 执行时候 没有任何处于等待
}
