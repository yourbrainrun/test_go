package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "123123A12312A2342566Ajjjj"
	sli := strings.Split(str, "A")

	fmt.Println(sli)
	fmt.Println(strings.Join(sli, ""))


	str1 := "qqqqqqqq"
	w, _ := os.Create("test.txt")

	rp := strings.NewReplacer(str,str1)
	n, _ := rp.WriteString(w, "gohomeDebug")
	fmt.Println(n)

	rp.Replace("dig")

	str123 := "123"
	str124 := "1234"
	fmt.Println(str124[len(str123):])
}
