package main

import (
	"fmt"
)

const (
	a1 = iota
	_  = iota
	a3 = iota
)

const (
	n1 = 10
	n2
	n3
)

func main() {
	fmt.Println(n2)
	fmt.Println(n3)
}
