package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`{"key":"value"}`)
	fmt.Println(buf.String())

	myMap := map[string]string{}
	_ = json.Unmarshal(buf.Bytes(), &myMap)

	fmt.Println(myMap["key"])
}
