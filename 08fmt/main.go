package main

import "fmt"

func main() {
	var n = 100
	//打印类型
	fmt.Printf("%T\n", n)
	// 打印值
	fmt.Printf("%v\n", n)
	// 打印二进制
	fmt.Printf("%b\n", n)
	// 打印十进制
	fmt.Printf("%d\n", n)
	// 打印八进制
	fmt.Printf("%o\n", n)
	// 打印十六进制
	fmt.Printf("%x\n", n)

	var s = "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}
