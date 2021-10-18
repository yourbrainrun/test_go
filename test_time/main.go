package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())

	parse, err := time.Parse("2006-01-02 15:04:05", "2021-08-25 14:25:33")
	if err != nil {
		return
	}

	fmt.Println(parse.Unix())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	cun, err := time.Parse("2006-01-02 15:04:05", "2022-02-01 00:00:01")

	hours := cun.Sub(time.Now()).Hours()

	fmt.Println(math.Floor(hours/24), " day ", int(hours)%24, " hours")

}
