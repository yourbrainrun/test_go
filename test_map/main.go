package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Info struct {
		Name string
		Age  int
		Desc string
	}
	var info Info
	info.Desc = "test"
	info.Age = 20
	info.Name = "xx"
	test(info)

	fmt.Println(map[string]string{"name": "test"})
}

func test(info interface{}) {
	bytesData, _ := json.Marshal(info)
	mapData := make(map[string]interface{}, 0)
	mapData2 := make(map[string]interface{}, 0)
	_ = json.Unmarshal(bytesData, &mapData)

	for key, value := range mapData {
		mapData2[key] = value
	}
	mapData2["other"] = 123

	bytesData2, _ := json.Marshal(mapData2)
	fmt.Println(string(bytesData2))
}

func test1() {
	maps := make(map[string]string, 0)
	maps["11"] = "123"

	if val, ok := maps["112"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("key error")
	}
}
