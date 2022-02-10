package main

import "fmt"

// make函数创造切片

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 不能用== nil来判断切片是否为空，要用len(s) == 0

	var ss1 []int
	ss2 := []int{}
	ss3 := make([]int, 0)
	fmt.Printf("ss1=%v len(ss1)=%d cap(ss1)=%d\n", ss1, len(ss1), cap(ss1))
	fmt.Printf("ss2=%v len(ss2)=%d cap(ss2)=%d\n", ss2, len(ss2), cap(ss2))
	fmt.Printf("ss3=%v len(ss3)=%d cap(ss3)=%d\n", ss3, len(ss3), cap(ss3))

	// 切片的遍历
	for i := 0; i < len(ss1); i++ {
		fmt.Println(ss1[i])

	}
}
