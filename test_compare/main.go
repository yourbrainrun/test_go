package main

import "fmt"

func test1() {
	fmt.Println("test1")
}
func test2() {
	fmt.Println("test2")
}

func main() {
	// Interface
	var inter1 interface{}
	var inter2 interface{}

	if inter2 == inter1 {
		fmt.Println("compare interface OK")
	}

	//if test1 == test2 {
	//}
	/*
		slice1 := make([]int, 1)
		slice2 := make([]int, 1)

		if slice2 == slice1 {
		}
	*/

	var int1 = 20
	var int2 = 30

	if &int1 == &int2 {
		fmt.Println("compare value pointer OK")
	} else {
		fmt.Println("compare value pointer OK")
	}

	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)
	if chan1 == chan2 {
		fmt.Println("compare channel OK")
	} else {
		fmt.Println("compare channel OK")
	}

	// 函数 切片 map 无法直接比较
	// 切片 map  函数 interface channel
}
