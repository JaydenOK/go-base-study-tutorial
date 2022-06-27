package main

import (
	"fmt"
)

func assert(i interface{}) {
	//断言值在最左边，第二个参数为断言成功与否Bool
	v, ok := i.(int) //获取int类型
	fmt.Println(v, ok)
}

//当给 assert 函数传递 Steven Paul 时，由于 i 的具体类型不是 int，ok 赋值为 false，而 v 赋值为 0（int 的零值）。该程序打印：
//
//56 true
//0 false
func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
