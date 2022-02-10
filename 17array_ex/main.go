package main

import "fmt"

func main() {

	arr1 := [...]int{4, 5, 6}

	sum := 0

	for _, v := range arr1 {
		sum = sum + v
	}

	fmt.Println(sum)

}
