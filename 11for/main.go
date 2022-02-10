package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// fmt.Println(i)
	}

	// 死循环
	// for {
	// 	fmt.Println("ss")
	// }

	s := "hello world"
	for i, v := range s {
		fmt.Printf("%d,%c\n", i, v)
	}

	s2 := "hello world"
	for _, v := range s2 {
		fmt.Printf("%c\n", v)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
