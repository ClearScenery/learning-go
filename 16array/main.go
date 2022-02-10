package main

import "fmt"

// 数组： var 变量名 [数组长度]数组类型

func main() {
	var a1 [3]bool // [true true false]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化：默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串：“”）
	fmt.Println(a1, a2)
	a1 = [3]bool{true, true, true}
	// 根据初始值自动推断数组长度
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a3)

	// 根据索引来初始化
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	// 根据索引遍历
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}

	// 2. for range遍历
	for _, v := range a4 {
		fmt.Println(v)
	}

	// 多维数组
	// [[1,2]]
	var a11 [3][2]int

	a11 = [3][2]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}

	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型

	b1 := [3]int{1, 2, 3}
	b2 := b1
	b1[0] = 100
	fmt.Println(b2)

	a22 := [...][2]int{
		{1, 2},
	}

	fmt.Println(a22)

	// 不支持
	// a33 := [2][...]int{
	// 	{1},
	// 	{2},
	// }
}
