package main

import (
	"fmt"
	"os"
)

func main() {
	foo()
}

func foo() {
	defer func() {
		if bar() != nil {
			os.Exit(0)
		}
	}()
	panic(3)
}

func bar() error {
	defer func() {
		if x := recover(); x != nil {
			fmt.Printf("bar is panicking\n")
			_, _ = x.(error)
		}
	}()
	if horribleNewProblemHappens() {
		fmt.Println("lowwer floor \n")
		panic("disaster\n")
	}

	fmt.Println(" second floor\n")
	return nil
}

func horribleNewProblemHappens() bool {
	return true
}
