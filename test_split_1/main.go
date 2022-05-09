package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	testSplit()
	return


	testPath := "test://livealihb.wangxiao.eaydu.com"
	OriginUrl := "livealihb.wangxiao.eaydu.com"

	urlDetail, err := url.Parse(testPath)
	if err != nil {
		return
	}

	pathSlice := strings.Split(testPath, "/")
	var retStr string
	if urlDetail.RawQuery == "" {
		retStr = fmt.Sprint("%s://%s/%s?cdn=wangsu/%s", "rtmp", OriginUrl, pathSlice[1], pathSlice[2])
	}
	retStr = fmt.Sprintf("%s://%s/%s?%s&cdn=wangsu/%s", "rtmp", OriginUrl, pathSlice[1], urlDetail.RawQuery, pathSlice[2])

	fmt.Println(retStr)
}

func testSplit() {
	OriginUrl := "livealihb.wangxiao.eaydu.com"
	sli := strings.Split(OriginUrl, ":")
	fmt.Println(sli)
	fmt.Println(sli[0])
}
