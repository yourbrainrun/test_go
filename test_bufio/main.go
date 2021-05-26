package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := "abcdefghijklmnopqrstuvwxyz1234567890"
	r := strings.NewReader(str)

	buf := bufio.NewReader(r)
	bytesinf, _ := buf.Peek(5)
	fmt.Println(string(bytesinf))
	fmt.Println(buf.ReadString('\n'))

	bytesinf, _ = buf.Peek(5)
	fmt.Println(string(bytesinf))

	buf.Reset(r)
	buf = bufio.NewReader(r)
	bytesinf, _ = buf.Peek(5)
	fmt.Println(string(bytesinf))  
}
