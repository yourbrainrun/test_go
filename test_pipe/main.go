package main

import (
	"fmt"
	"os"
)

func main() {
	if pr, pw, err := os.Pipe(); err == nil {

		m, _ := pw.WriteString("sdfasdfasdf")
		fmt.Println(m)

		var bb = make([]byte, 10)
		n, _ := pr.Read(bb)

		fmt.Println(n)
		fmt.Println(string(bb))

	}
}
