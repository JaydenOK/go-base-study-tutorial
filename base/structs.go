package main

import (
	"fmt"
)

//结构体的声明
//type Employee struct {
//	firstName string
//	lastName  string
//	age       int
//}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type User struct {
	name string
	age  int
}

//结构体方法
func (user User) getName() string {
	return user.name
}

func niMing() {
	//创建匿名结构体
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	//匿名字段结构体，按顺序赋值, 默认字段类型 即为字段名
	type Person struct {
		string
		int
	}

	p := Person{
		string: "hello",
		int:    0,
	}
	fmt.Println(p)

}

//提升字段:，Person 结构体有一个匿名字段 Address，而 Address 是一个结构体。现在结构体 Address 有 city 和 state 两个字段，访问这两个字段就像在
// Person 里直接声明的一样，p.city访问
type Address struct {
	city, state string
}

//提升方法，Address方法，Person2也可直接访问
func (a Address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type Person2 struct {
	name string
	age  int
	Address
}

//当一个函数有一个值参数，它只能接受一个值参数。
//当一个方法有一个值接收器，它可以接受值接收器和指针接收器

func main() {
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	user := User{
		name: "李白",
		age:  43,
	}

	fmt.Println(user.getName())

	////////////////////////////////////////////
	//在上面代码中，结构体类型 image 包含一个 map 类型的字段。由于 map 类型是不可比较的，因此 image1 和 image2 也不可比较。如果运行该程序，
	// 编译器会报错：main.go:18: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)。
	type image struct {
		data map[int]int
	}
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}

	fmt.Println(image1, image2)
	//if image1 == image2 {
	//	fmt.Println("image1 and image2 are equal")
	//}

}
