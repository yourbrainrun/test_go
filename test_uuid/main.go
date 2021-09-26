package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// panic on error
	for i := 0; i < 5; i++ {
		u2 := uuid.NewV4()
		fmt.Println(fmt.Sprintf("%s", u2))
		fmt.Println(u2.String())
	}
}
