package main

import "fmt"

type User2 struct {
}
type MyUser1 User2
type MyUser2 = User2

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User2) m2() {
	fmt.Println("User.m2")
}

func main() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}
