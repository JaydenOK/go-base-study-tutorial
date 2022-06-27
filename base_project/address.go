package main

import "fmt"

func main() {
	b := "111"
	var a *string = &b                 //& 操作符用于获取变量的地址，或者这样写：a := &b
	fmt.Printf("Type of a is %T\n", a) //类型： *string
	fmt.Println(a)                     //0xc0000301f0
	fmt.Println(*a)                    //111 	//指针的解引用

	*a = "222" // 把a指向的值(b)改为 222

	fmt.Println(*a)
	fmt.Println(b)

}
