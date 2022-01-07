package main

import "fmt"

func main() {
	test()
}

func test() {
	aa := [...]int{
		1, 2, 3, 4, 5,
	}
	bb := [...]int{
		6, 7, 8,
	}

	//cc := append(aa, bb...)
	//fmt.Println(cc)
	fmt.Println(aa, bb)
}
