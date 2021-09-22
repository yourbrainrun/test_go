package main

import (
	"fmt"
	"strconv"
)

func main() {
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
