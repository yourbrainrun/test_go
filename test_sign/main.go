package main

import (
	"fmt"
	"github.com/yourbrainrun/test_go/test_sign/roc"
)

func main() {
	client := roc.Client{
		AuthKey:   "9pkoxvfHy6HgANuX",
		SecretKey: "dasfsfa135135513",
	}
	var expiredAt int64
	expiredAt = 1655885187
	authToken, err := client.GetToken("11200", "0", expiredAt)
	if err != nil {
		panic(err)
	}

	fmt.Println(authToken)
}
