package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	str := "test"

	refType := reflect.TypeOf(str)

	newType := reflect.New(refType)
	fmt.Println(newType)

	fmt.Println(strings.Split("", ","))
}
