package main

import (
	"fmt"
	"github.com/yourbrainrun/test_go/test_compare_struct/model"
)

func main() {
	var pro1 model.Project
	Set(&pro1)
	fmt.Println(pro1.IsEmpty(), "pro1")

	var pro2 model.Project
	fmt.Println(pro2.IsEmpty(), "pro2")
}

func Set(pointer *model.Project) {
	// 注意 使用指针 和 给指针变量赋值 给指针对应的变量属性赋值 完全不同
	pointer.Name = "hua hua"
}
