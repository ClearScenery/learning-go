package main

import "fmt"

func main() {
	// 基于数据定义切片
	a := [5]int{1, 2, 3, 4, 5}
	// 定义切片 a[1],a[2],a[3]
	b := a[1:4]
	// [2 3 4]
	fmt.Println(b)
	// type of b:[]int
	fmt.Printf("type of b:%T\n", b)

	a1 := [...]int{1, 2, 3, 4, 5}
	s1 := a1[:3]
	s2 := a1[3:]
	s3 := a1[:]

	fmt.Println(s1, s2, s3)

	a2 := [...]int{2, 3, 4, 5}
	ss1 := a2[1:3]
	fmt.Println(ss1)

	ss1[0] = 5
	fmt.Println(a2)

	// 切片是引用类型，都指向了地层的同一个数组
	fmt.Println(len(ss1), cap(ss1))

}

func arraySum(arr [3]int) int {
	sum := 0

	for _, v := range arr {
		sum = sum + v
	}

	return sum
}
