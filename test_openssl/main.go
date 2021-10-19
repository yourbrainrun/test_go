package main

import (
	"encoding/base64"
	"fmt"
	"github.com/yourbrainrun/test_go/test_openssl/crypt"
)

func main() {
	key := "lOLWf2ubswSDrUSt6lfwLQ"

	svc1 := crypt.NewOpensslAES128ECB(key)
	enStr := "Yzyz8R6AmgHvLQv5JoQq1ivFg7PXfg/ygNGhA1cjY/tqMVtejCoQM0OZTeF1bS4Z"
	decodeString, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))
	fmt.Println(len(key))
	str1 := svc1.Decrypt(string(decodeString))
	fmt.Println(str1 + "|")

	return
}
