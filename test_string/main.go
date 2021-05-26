package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123123A12312A2342566Ajjjj"
	sli := strings.Split(str, "A")

	fmt.Println(sli)
	fmt.Println(strings.Join(sli, ""))
}
