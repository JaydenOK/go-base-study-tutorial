package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
}

func rectProps2(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
