package tests

import (
	"fmt"
	"testing"
)

//值得注意的是，因为 map 是无序的，因此对于程序的每次执行，不能保证使用 `for range` 遍历 map 的顺序总是一致的，而且遍历的顺序也不完全与元素添加的顺序一致。
//定义初始化时，逗号结尾
func TestMap(t *testing.T) {
	//salary := make(map[string]int)   //定义，使用零值
	var m = map[string]string{
		"a": "aa",
		"b": "bb",
	}
	m["c"] = "cccc"
	delete(m, "a")
	fmt.Println(m)
	fmt.Println("###################", "")

	m2 := map[string]map[string]string{
		"a": {
			"aa": "aaa",
		},
		"b": {
			"bb": "bbb",
		},
	}
	m2["d"] = map[string]string{
		"dd": "ddddd",
	}
	for key, value := range m2 {
		for k, v := range value {
			fmt.Println("key:"+key, k, v)
		}
	}
}
