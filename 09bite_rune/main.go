package main

import "fmt"

func main() {
	// rune类型
	a := '中'
	// byte类型
	b := 'x'

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转换成rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := byte('H') // byte(unit8)

	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Printf("%T %v\n", f, f)

}
