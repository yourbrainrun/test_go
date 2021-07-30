package main

import (
	"fmt"
	"github.com/yourbrainrun/test_go/test_rwmutex/MyMap"
	"time"
)

func main() {

	temp := make(chan int)
	data := MyMap.NewShareData()

	go func() {
		for i := 0; i < 10; i++ {
			ret := data.Get("test", 1000)
			fmt.Println(ret, " thread 1")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ret := data.Get("test", 1)
			ret2 := data.Get("foo", 1)

			fmt.Println(ret, "|", ret2, " thread 2")
		}
	}()
	time.Sleep(10 * time.Millisecond)
	go func() {
		_ = data.Set("foo", "bar")
	}()
	<-temp
}
