package types

import (
	"fmt"
	"testing"
)

type a = string
type b string

func SayA(str a) {
	fmt.Println(str)
}

func SayB(str b) {
	fmt.Println(str)
}

func TestTypeAlias(t *testing.T) {
	var str = "test"
	SayA(str)

	//错误参数传递，str是字符串类型，不能赋值给b类型变量
	//SayB(str)
}
