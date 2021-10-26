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

	tt, err := time.Parse(time.RFC3339, "2021-10-26T19:00:55+08:00")
	if err != nil {
		return
	}
	fmt.Println(tt.Unix(), "rfc3339", time.Now().Unix())

	cun, err := time.Parse("2006-01-02 15:04:05", "2022-02-01 00:00:01")

	hours := cun.Sub(time.Now()).Hours()

	fmt.Println(math.Floor(hours/24), " day ", int(hours)%24, " hours")

	parse, err = time.Parse("2006-01-02 15:04:05", "2021-10-26 18:12:56")
	if err != nil {
		return
	}
	if parse.Unix()+86400 < time.Now().Unix() {
		fmt.Println("guo qi ")
	} else {
		fmt.Println("0000000 ")
	}
}
