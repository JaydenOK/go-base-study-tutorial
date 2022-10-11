package asserts

import (
	"fmt"
)

//类型断言就是将接口类型的值(x)，转换成类型(T)。格式为：x.(T)；
//类型断言的必要条件就x是接口类型，非接口类型的x不能做类型断言；
//T可以是非接口类型，如果想断言合法，则T必须实现x的接口；
//T也可以是接口，则x的动态类型也应该是接口T；
//类型断言如果非法，运行时会导致错误，为了避免这种错误，应该总是使用下面的方式来进行类型断言

// 需要注意的如果不接收第二个参数也就是ok，这里失败的话则会直接panic这里还存在一种情况就是x为nil同样会panic
func assert(i interface{}) {
	//断言值在最左边，第二个参数为断言成功与否Bool
	v, ok := i.(int) //获取int类型
	fmt.Println(v, ok)
}

// 当给 assert 函数传递 Steven Paul 时，由于 i 的具体类型不是 int，ok 赋值为 false，而 v 赋值为 0（int 的零值）。该程序打印：
//
// 56 true
// 0 false
func main() {
	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
