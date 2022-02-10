package main

import (
	"fmt"
	"strings"
)

func main() {
	// int32
	c1 := 'a'
	c2 := '中'
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)

	s := "I'm ok"
	fmt.Println(s)

	s2 := `d:\demo\test01`
	fmt.Println(s2)
	fmt.Println(len(s2))

	s3 := "hello"
	s4 := ",world"
	fmt.Printf("%s%s\n", s3, s4)
	ss1 := fmt.Sprintf("%s%s", s3, s4)
	fmt.Println(ss1)

	ss2 := strings.Split(ss1, ",")
	fmt.Println(ss2)

	fmt.Println(strings.Contains(ss1, "hello"))
	fmt.Println(strings.HasPrefix(ss1, "hello"))
	fmt.Println(strings.HasSuffix(ss1, "hello"))

	s5 := "abcdeb"
	fmt.Println(strings.Index(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "b"))

	fmt.Println(strings.Join(ss2, "+"))
}
