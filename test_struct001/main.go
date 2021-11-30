package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Teacher struct {
	Class string
	User
}

func main() {
	// 自动做类型转换适配
	T1 := &Teacher{
		Class: "1111",
		// Name 无法直接使用
		User: User{
			Name: "china",
			Age:  5000,
		},
	}

	fmt.Println(T1)

	var T2 Teacher
	T2.Class = "optimize"
	T2.Name = "dog dan"

	fmt.Println(T2)
}
