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

	_, _ = buf.Discard(10)
	fmt.Println(string(bytesinf))
	fmt.Println(buf.ReadString('\n'))

	buf.Reset(r)
	buf = bufio.NewReader(r)
	bytesinf, _ = buf.Peek(5)
	fmt.Println(string(bytesinf))
}
