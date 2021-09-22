package main

import (
	"encoding/base64"
	"fmt"
	"github.com/yourbrainrun/test_go/test_openssl/crypt"
)

func main() {
	key := "testkey"

	svc1 := crypt.NewOpensslAES128ECB(key)
	enStr := "xquWr9JWshGZQSmcNuiuiKsHsI1SrJ27qV78RBBvXplqMVtejCoQM0OZTeF1bS4Z"
	decodeString, err := base64.StdEncoding.DecodeString(enStr)
	if err != nil {
		return
	}
	fmt.Println(string(decodeString))
	fmt.Println(len(key))
	str1 := svc1.Decrypt(string(decodeString))
	fmt.Println(str1)

	return
}
