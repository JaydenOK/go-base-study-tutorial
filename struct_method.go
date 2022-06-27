package main

import (
	"fmt"
)

type Employee2 struct {
	name string
	age  int
}

/*
使用值接收器的方法。
*/
func (e Employee2) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *Employee2) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee2{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//使用指针接收器的方法。能改变传入的变量,
	//由于 changeAge 方法是使用指针 (e *Employee) 接收器的，所以在调用 (&e).changeAge(51) 方法对 age 字段做出的改变对调用者将是可见的。
	// (&e).changeAge(51)
	e.changeAge(51) //简写, 改变调用者e
	fmt.Printf("\nEmployee age after change: %d", e.age)

}
