package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	map1 := make(map[string]interface{})
	map1["ab"] = "ab1"
	map1["ab2"] = "ab21"
	map1["ab3"] = "ab31"

	map2 := make(map[string]interface{})
	str, _ := json.Marshal(map1)
	json.Unmarshal(str, &map2)

	map2["test"] = "3test"
	fmt.Println(fmt.Sprintf("map1:%p : %v", map1, map1))
	fmt.Println(fmt.Sprintf("map2:%p : %v", map2, map2))

	str1 := "123123"
	str2 := "3333"
	var buf bytes.Buffer
	buf.WriteString(str1)
	str2 = buf.String()
	fmt.Println(str2)
}
