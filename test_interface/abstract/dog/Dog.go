package dog

import "fmt"

type Dog struct {
	Name string
	Age  uint
	Addr string
}

func (dog Dog) Call() {
	fmt.Println("dog ", dog.Name, "eat  and age ", dog.Age)
}

func GetDog() Dog {
	return Dog{Name: "dog ", Age: 1, Addr: "bj"}
}
