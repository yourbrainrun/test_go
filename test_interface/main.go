package main

import "fmt"

type sun interface {
	say()
}

type holy interface {
	sun
	bible()
}

type My struct {
	Name string
}

func (me My) say() {
	fmt.Println(me.Name)
}

func (me *My) bible() {
	fmt.Println(me.Name, " body not say not")
}

func main() {
	var t My
	t.Name = "dog testis"

	var t1 = &My{}
	t1.Name = "xi wa dog testis"

	ctrl(t)
	holyBible(t1)
}

func ctrl(fac sun) {
	fac.say()
}

func holyBible(ho holy) {
	ho.say()
	ho.bible()
}
