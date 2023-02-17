package maps

import (
	"encoding/json"
	"fmt"
)

//检查key是否存在
func InMapString(m map[string]string, key string) bool {
	if _, isExist := m[key]; isExist {
		//参数2:true，参数1为其对应值
		return true
	}
	return false
}

//检查key是否存在(整型)
func InMapInt(m map[int]int, key int) bool {
	if _, isExist := m[key]; isExist {
		return true
	}
	return false
}

//合并，键相同，后面的m2覆盖前面m1的值
func Merge(m1 map[string]string, m2 map[string]string) map[string]string {
	for key, value := range m2 {
		m1[key] = value
	}
	return m1
}

// 删除元素，使用内置函数
func Unset(m map[string]string, key string) {
	delete(m, key)
}

// map是无序的，可放入切片排序，输出
func SortNotUseful(m map[string]string) {
}

// map转json字符串
func Stringify(m map[string]string) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(m); err != nil {
		panic(interface{}(err))
	} else {
		return string(bytes)
	}
}

// json字符串转map（解析）
func Parse(s string) map[string]interface{} {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		panic(interface{}(err))
	}
	return m
}

func JsonEncode(m map[string]string) string {
	return Stringify(m)
}

func JsonDecode(s string) map[string]interface{} {
	return Parse(s)
}

//解析json字符，按层级获取数据，先断言再获取
func decodeTest(s string) {
	str := `{"a": "11", "b": "22", "c": "33", "d": "666", "user": {"name":"LiMing","age":12,"tags":["a","b"]}}`
	m5 := JsonDecode(str)
	user := m5["user"].(map[string]interface{})
	fmt.Println(user["tags"])
	tags := user["tags"].([]interface{})
	fmt.Println(tags[0])
}
