package main

import "fmt"

func main() {

	// age := 19
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	} else {
		fmt.Println("该写作业了")
	}

	// age作用域
	if age := 19; age > 35 {
		fmt.Println("人到中年")
		fmt.Println(age)
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("少年")
	}

	// fmt.Println(age)
}
