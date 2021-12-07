package main

import "fmt"

func main() {
	test()
}

func test() {
	x:=1
	for i := 0; i < 100; i++ {
		x/=x+1
		fmt.Println(x)
	}
}
