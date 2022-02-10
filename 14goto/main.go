package main

import "fmt"

func main() {
	// break跳出一层层循环
	for i := 0; i < 10; i++ {
		if i == 2 {
			goto flag
		}
		fmt.Printf("%d\n", i)
	}
flag:
	fmt.Println("goto here")
}
