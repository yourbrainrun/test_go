package main

import "fmt"

type Man interface {
	Get()
	Set(string)
}
type User struct {
	Name string
	Age  int
}

func (u User) Get() string {
	return u.Name
}

func (u *User) Set(name string) {
	u.Name = name
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func main() {
	u1 := NewUser("lili", 23)
	fmt.Println(u1.Get())
	u1.Set("lili men")
	fmt.Println(u1.Get())

	u2 := &User{
		Name: "lilei",
		Age:  12,
	}
	fmt.Println(u2.Get())
	u2.Set("li dong men")
	fmt.Println(u2.Get())

	// 自动做类型转换适配
}
