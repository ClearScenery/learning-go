package main

import "fmt"

// 默认值false
var b1 bool

func main() {
	fmt.Println(b1)

	i1 := 1
	i1 += 1

	fmt.Printf("%T value:%v", b1, b1)
}
