package main

import "fmt"

type MyData struct {
	Name string
	Age  int
}

func main() {
	var my MyData

	fmt.Println(fmt.Sprintf("%p", &my))

	testP(&my)

}
func testP(aa *MyData) {

	fmt.Println(fmt.Sprintf("%p", &aa), "param")
	fmt.Println(fmt.Sprintf("%p", aa), "param pointer")

	testPP(aa)

	testPPP(&(*aa))

	fmt.Println(aa.Age)
}

func testPP(aa *MyData) {

	fmt.Println(fmt.Sprintf("%p", &aa), "2 param")
	fmt.Println(fmt.Sprintf("%p", aa), "2 param pointer")

	aa.Age = 22
}

func testPPP(aa *MyData) {

	fmt.Println(fmt.Sprintf("%p", &aa), "3 param")
	fmt.Println(fmt.Sprintf("%p", aa), "3 param pointer")
}
