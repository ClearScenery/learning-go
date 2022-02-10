package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	// 调用append函数必须用原来的切片变量接收返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := []string{"西安", "成都"}
	// ...展开切片
	s1 = append(s1, s2...)
}
