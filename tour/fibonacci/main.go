package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	i := 0
	j := 1
	return func() int {
		res := i
		i = j
		j = res + j
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
