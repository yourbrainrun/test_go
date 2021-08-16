package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

	r.Reset(str)
	buf = bufio.NewReader(r)
	bytesinf, _ = buf.Peek(5)
	fmt.Println(string(bytesinf))

	strToFile()
	strFile()
}

func strToFile() {
	str := "test hello "
	ioR := strings.NewReader(str)
	var fi *os.File
	var err error
	fileName := "test_bufio/test.txt"
	if fi, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644); err == nil {
	} else {
		fmt.Println(" open ")
		fi, _ = os.Create(fileName)
	}

	buf := bufio.NewReader(ioR)
	buf.WriteTo(fi)

	fi.Close()
}

func strFile() {
	str := "test hello "
	ioR := strings.NewReader(str)
	var fi *os.File
	var err error
	fileName := "test_bufio/test1.txt"
	if fi, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0644); err == nil {
	} else {
		fmt.Println(" open ")
		fi, _ = os.Create(fileName)
	}

	io.Copy(fi, ioR)

	fi.Close()
}
