package main

import (
	"fmt"
	"github.com/yourbrainrun/test_go/test_lua/one"
	"github.com/yourbrainrun/test_go/test_lua/two"
)

func main() {
	one.ExecLuaOne()
	fmt.Println("=-----------------=")
	two.ExecTwo()
}
