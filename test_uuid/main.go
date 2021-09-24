package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// panic on error
	u2 := uuid.NewV4()
	fmt.Println(u2)
}
