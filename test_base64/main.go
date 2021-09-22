package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "test123 touch you head and say ok"
	enCode := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(enCode)

	decodeString, err := base64.StdEncoding.DecodeString(enCode)
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))
}
