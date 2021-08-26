package main

import (
	"fmt"
	"time"
)

func main() {
	parse, err := time.Parse("2006-01-02 15:04:05", "2021-08-25 14:25:33")
	if err != nil {
		return
	}
	diff := time.Now().Sub(parse).Seconds()

	fmt.Println(diff)
	fmt.Println(time.Now().Unix(), "unix")
	fmt.Println(parse.Sub(time.Now()).Seconds())
}
