package main

import (
	"fmt"
	"github.com/forgoer/openssl"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//test1()
	//test2()

	//fmt.Println(test3())
	//fmt.Println(test4())
	//func5()
	split()
}

func split() {
	str := "rtmp://push.roadofcloud.net/live/b502a06f034940039a6abae62f639001"
	str = "1"
	data := strings.Split(str, "/")
	fmt.Println(data)
	fmt.Println(data[len(data)-1])
}

func func5() {
	str := "http://test.baidu.com/mix/test"
	inNum := strings.Index(str, "/mix/")
	fmt.Println(inNum)
}

func test4() bool {
	num, _ := strconv.Atoi("")
	fmt.Println(num)
	return true

	str := "hello_"
	length := len(str)
	if length > 1 && length <= 65 {
		match := regexp.MustCompile(`[\s]+`)
		ret := match.FindAllString(str, -1)

		length = len(ret)
		fmt.Println(ret)
		if length > 0 {
			return true
		}

		return false
	} else {
		return false
	}
}

func test3() bool {
	str := "hello_"
	length := len(str)
	if length > 1 && length <= 65 {
		match := regexp.MustCompile(`^[\w]+$`)
		ret := match.FindAllString(str, -1)

		length = len(ret)
		if length > 0 {
			return true
		}

		return false
	} else {
		return false
	}
}

func test2() {
	str := "PG12EEPGEcOtjBOfmkrL4oHSGFYbjQ9R\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010"
	over := strings.TrimRight(str, openssl.ZEROS_PADDING)
	fmt.Println(over, "|")

	str2 := "PG12EEPGEcOtjBOfmkrL4oHSGFYbjQ9R "
	over1 := strings.TrimRight(str2, " ")
	fmt.Println(over1, "|")

	reg := regexp.MustCompile("\\s+")
	str3 := reg.ReplaceAllString(str, "")
	fmt.Println(str3, "|")
}

func test1() {
	str := "123123A12312A2342566Ajjjj"
	sli := strings.Split(str, "A")

	fmt.Println(sli)
	fmt.Println(strings.Join(sli, ""))

	str1 := "qqqqqqqq"
	w, _ := os.Create("test.txt")

	rp := strings.NewReplacer(str, str1)
	n, _ := rp.WriteString(w, "gohomeDebug")
	fmt.Println(n)

	rp.Replace("dig")

	str123 := "123"
	str124 := "1234"
	fmt.Println(str124[len(str123):])

	sli = strings.Split("sdn://wss.wangxiao.eaydu.com/live_ali/x_3_1119892_51266", "sdn://")

	fmt.Println(sli[0], "one")
	fmt.Println(sli[1])

	fmt.Println(strings.TrimRight("asdfasdfa    ", " ") + "|")
}
