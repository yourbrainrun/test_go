package main

import (
	"fmt"
	"time"
)

func main() {
	work()
}

func work() {
	//http://127.0.0.1:999/work

	fmt.Println("start work... ")

	go sonLine()
	fmt.Println("work over")
}

func sonLine() {
	fmt.Println("son line start run!")
	time.Sleep(time.Second * 5)
	fmt.Println("son line end run!")
}
