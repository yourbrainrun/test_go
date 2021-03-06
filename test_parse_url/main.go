package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "cdn://wss.wangxiao.eaydu.com/live_ali/x_3_test?key=d0e5159c95e4bd63c93ad467bb14f134&tm=74841c9c"
	//str = "cdn://wss.wangxiao.eaydu.com/live_ali/x_3_test"
	detail, _ := url.Parse(str)
	fmt.Println(detail.Host)
	fmt.Println(detail.Path)
	fmt.Println(detail.RawQuery)
	if detail.Port() == "" {
		fmt.Println("port empty")
	} else {
		fmt.Println(detail.Port(), "port")
	}

	query, err := url.ParseQuery(detail.RawQuery)
	if err != nil {
		return
	}
	for keys, values := range query {
		fmt.Println(keys, values[0], "||")
	}
}
