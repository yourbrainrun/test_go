package main

import "fmt"

func main() {
	f(true)
}

func f(doponic bool) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("panicking with value %v\n", x)
			//panic(x) // 回到恐慌状态
		}
		fmt.Printf("函数正常返回\n")
	}()
	fmt.Printf("之前\n")
	p(doponic)
	fmt.Printf("after\n")
}

func p(doponic bool) {
	if doponic {
		panic(3)
	}
}
