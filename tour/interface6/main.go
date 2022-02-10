package main

import "fmt"

type I interface {
	M()
}

type A struct {
	ap int
}

func (a *A) M() {
	fmt.Println(a.ap)
}

func main() {
	var i I
	var a *A
	i = a
	describe(i)
	i = &A{3}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
