package structs

import (
	"fmt"
	"testing"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 嵌套，扩展结构体方法
type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	//这个嵌入式接口允许Reverse使用另一个接口实现的方法
	Interface
}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	return &reverse{data}
}

type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x StringSlice) Sort()              { fmt.Println("aaa") }

// 嵌入
func TestEmbed(t *testing.T) {

}
