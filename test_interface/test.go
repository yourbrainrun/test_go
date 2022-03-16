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

func ctrl(fac sun) {
	fac.say()
}

func holyBible(ho holy) {
	ho.say()
	ho.bible()
}
