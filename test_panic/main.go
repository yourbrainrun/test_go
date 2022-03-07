package main

import "fmt"

func main() {
	f(true)
}

func f(doPanic bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("panicking with value %v\n", x)
			//panic(x) // 回到恐慌状态
		}
		fmt.Printf("函数正常返回\n")
	}()
	fmt.Printf("之前\n")
	p(doPanic)
	fmt.Printf("after\n")
}

func p(doPanic bool) {
	if doPanic {
		panic(3)
	}
}
