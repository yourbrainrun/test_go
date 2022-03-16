package main

import (
	"github.com/yourbrainrun/test_go/test_interface/abstract"
	"github.com/yourbrainrun/test_go/test_interface/abstract/dog"
	"github.com/yourbrainrun/test_go/test_interface/abstract/pig"
)

func main() {
	//test1()
	test2()
}

func test1() {
	var t My
	t.Name = "dog testis"

	var t1 = &My{}
	t1.Name = "xi wa dog testis"

	ctrl(t)
	holyBible(t1)
}

func test2() {
	dog1 := dog.GetDog()
	factory(dog1)

	pig1 := pig.GetPig()
	factory(pig1)
}

func factory(animal abstract.Action) {
	animal.Call()
}
