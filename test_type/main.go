package main

import "fmt"

func main() {
	my := map[string]interface{}{}
	my["test"] = 1212
	my["test1"] = "string"

	for i, v := range my {
		fmt.Println(i)
		switch val := v.(type) {
		case string:
			fmt.Println(v, "string")
		case int:
			fmt.Println(v, "int")
		default:
			fmt.Println(val)
		}
	}
}
