package pig

import "fmt"

type Pig struct {
	Name string
	Age  uint
	Addr string
}

func (pig Pig) Call() {
	fmt.Println("pig ", pig.Name, "eat  and age ", pig.Age)
}

func GetPig() Pig {
	return Pig{Name: "dog ", Age: 1, Addr: "bj"}
}
