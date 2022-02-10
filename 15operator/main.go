package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// 单独的语句，不能放在等号的右边
	a--
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a < b) // Go语言是强类型，相同类型才能比较
	fmt.Println(a > b)
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	age := 12

	// 逻辑运算符
	if age >= 18 && age <= 60 {
		fmt.Println("苦逼上班")
	} else {
		fmt.Println("不用上班")
	}

	if age <= 18 || age >= 60 {

	}

	isMarked := false
	fmt.Println(!isMarked)

	// 位运算符
	var x int
	x = 10
	x += 1
	x -= 1
	fmt.Println(x)

}
