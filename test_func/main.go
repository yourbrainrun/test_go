package main

import "fmt"

func main() {
	str := create(sign)
	fmt.Println(str)

	str = create(nil)
	fmt.Println(str)
}

func create(signFunc func(map[string]string) string) string {
	li := make(map[string]string, 0)
	li["sort"] = " test ok"

	if signFunc == nil {
		return "nil sign func"
	}
	return signFunc(li)
}

func sign(sli map[string]string) string {
	var str string
	for key, value := range sli {
		str = str + key + value
	}

	return str
}
