package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	var dd float64
	dd = 23.22
	ss := fmt.Sprintf("%f", dd)
	fmt.Println(ss)

	test := sync.Pool{
		New: func() interface{} {
			return rand.Int()
		},
	}

	val := test.Get().(int)
	fmt.Println(val)

	val = test.Get().(int)
	fmt.Println(val)

	test.Put(12)
	fmt.Println(test.Get().(int))

	test.Put(15)
	runtime.GC() // 有时候可以出现 15 有时候是 0 不稳定
	fmt.Println(test.Get().(int))
}
