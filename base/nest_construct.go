package main

import (
	"fmt"
)

//嵌套结构体，外层结构体可直接访问，嵌套结构体属性，方法

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type poster struct {
	title   string
	content string
	//嵌套结构体，外层结构体可直接访问，嵌套结构体的属性、方法
	author
}

func (p poster) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

//切片保存多个（使用数组则为固定长度）
type website struct {
	posts []poster
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := poster{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := poster{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := poster{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []poster{post1, post2, post3},
	}
	w.contents()
}
