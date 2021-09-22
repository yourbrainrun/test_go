package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())

	parse, err := time.Parse("2006-01-02 15:04:05", "2021-08-25 14:25:33")
	if err != nil {
		return
	}

	fmt.Println(parse.Unix())
}
