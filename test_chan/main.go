package main

import (
	"fmt"
	"time"
)

func main() {
	sendIntervalSecond := time.Second * 1
	recvIntervalSecond := time.Second * 2
	intChan := make(chan int)
	go func() {
		var t0, t1 int64
		for i := 0; i < 5; i++ {
			intChan <- i
			time.Sleep(sendIntervalSecond)
			t1 = time.Now().Unix()
			if t0 == 0 {
				fmt.Println("send %d:", i)
			} else {
				fmt.Println("time: %d data:%d", t1-t0, i)
			}
			t0 = time.Now().Unix()
		}
		close(intChan)
	}()
Loop:
	for {
		var t0, t1 int64
		select {
		case data, ok := <-intChan:
			if !ok {
				break Loop
			}
			time.Sleep(recvIntervalSecond)
			t1 = time.Now().Unix()
			if t0 == 0 {
				fmt.Println("send %d:", data)
			} else {
				fmt.Println("time: %d data:%d", t1-t0, data)
			}
			t0 = time.Now().Unix()
		}
	}
}
