package main

import "fmt"

func main() {
	var n = 5
	if n == 1 {
		fmt.Println("大拇指")
	} else {
		// ..
	}

	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println(n)
	}

	switch n := 3; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("是奇数")
		// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
		fallthrough
	case 6:
		fmt.Println("fallthrough")
	default:
		fmt.Println("是偶数")
	}
}
