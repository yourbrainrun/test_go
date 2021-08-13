package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
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
