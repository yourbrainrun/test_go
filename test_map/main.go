package main

import "fmt"

func main() {
	maps := make(map[string]string, 0)
	maps["11"] = "123"

	if val, ok := maps["112"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("key error")
	}
}
