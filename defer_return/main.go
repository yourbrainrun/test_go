package main

import "fmt"

func main() {
	n := numDefer()
	fmt.Println(n)
}

func numDefer() (n int) {
	for i := 0; i < 100; i++ {
		defer func() {
			n++
		}()
	}
	return
}
