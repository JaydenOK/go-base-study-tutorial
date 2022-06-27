package main

import (
	"fmt"
)

//高阶函数
func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func main() {
	s := simple()
	fmt.Println(s(60, 7))
}
