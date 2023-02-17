package types

import (
	"base_project/prints"
	"fmt"
	"testing"
)

//我们可以使用type Name Type去自定义一个全新的数据类型。这个类型的变量就使用Name来声明。通常type用于将一个复杂的数据类型定义为一个全新的数据类型，便于后续使用。
//场景：扩展字符串切片，并为其增加扩展方法show()，实现MyTypeInterface方法

// 定义新接口
type MyTypeInterface interface {
	setName(string2 string)
	getName() string
}

type StringList []string

func (s StringList) setName(string2 string) {
	fmt.Println("setName")
}
func (s StringList) getName() string {
	fmt.Println("getName")
	return ""
}

func show(v MyTypeInterface) {
	prints.PrintVar(v)
}

func TestT(t *testing.T) {
	str := []string{"test1", "test2", "2222"}
	//需要调用show方法，处理字符串切片逻辑，需要转换成MyTypeInterface类型，而StringList基本类型为[]string，并且实现了MyTypeInterface类型
	list := StringList(str)
	show(list)

	//倒序源码实例
	//sort.Sort(sort.Reverse(sort.StringSlice(str))) //Reverse()函数所需的参数类型Interface, StringSlice实现了type MyTypeInterface interface{}，故str先强转Interface类型
}
