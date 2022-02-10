package main

import "fmt"

func main() {
	s1 := 101
	fmt.Printf("%d\n", s1)
	fmt.Printf("%b\n", s1)
	s2 := 0b101
	fmt.Printf("%b\n", s2)
	var s3 int = 0xff
	fmt.Printf("%x\n", s3)

	var s4 int = 0777
	fmt.Printf("%d\n", s4)

	var s5 int32 = 33
	fmt.Printf("%d\n", s5)
	fmt.Printf("%T\n", s5)

	s5 = int32(s4)
}
