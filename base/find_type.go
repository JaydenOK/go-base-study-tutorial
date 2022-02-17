package main

import "fmt"

type Describer interface {
	Describe()
}

type Printer interface {
	print()
}

type Person struct {
	name string
	age  int
}

//其原因是：对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。但接口中存储的具体值（Concrete Value）并不能取到地址，
// 因此在第 45 行，对于编译器无法自动获取 a 的地址，于是程序报错。
//实现Describer接口
func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

//实现Printer接口
func (p Person) print() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	//是不是一个接口类型 (is instance of Describer)
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}
