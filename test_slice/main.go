package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func main() {
	test2()
	// map slice 都是引用 在函数中修改会反馈到 函数调用处
	sli := make([]string, 5)
	testSlice(sli, true)
	fmt.Println(sli)

	mp := make(map[string]string, 5)
	testMap(mp)
	fmt.Println(mp)

	// slice append more then cap ,return no change
	testSlice(sli, false)
	fmt.Println(sli)
}

func test2() {
	tempLicense := make([]map[string]string, 1)
	tempLicense[0] = map[string]string{
		"token":   url.QueryEscape("122222332323"),
		"app_key": "23234",
	}

	fmt.Println(tempLicense)
}

func testSlice(temp []string, flag bool) {
	if flag {
		for i, _ := range temp {
			temp[i] = "slice" + fmt.Sprintf("%d", i)
		}
	} else {
		for i, _ := range temp {
			temp[i] = "slice" + fmt.Sprintf("%d", i)
		}
		temp = append(temp, "test append one ")
		fmt.Println(temp, "append one")
	}

}

func testMap(temp map[string]string) {
	fmt.Println(temp, "11")

	for i := 0; i < 5; i++ {
		temp["map"+strconv.Itoa(i)] = "mp value" + strconv.Itoa(i)
	}
}
