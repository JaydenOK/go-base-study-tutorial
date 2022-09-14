package convert

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_string_int_slice_byte_test(t *testing.T) {
	//ByteToString()
	AssertInterface()
}

// 字符串转Byte
func StringToByte(str string) []byte {
	b := []byte(str)
	return b
}

// byte转字符串
func ByteToString() {
	var str = "abc123"
	b := StringToByte(str)
	str2 := string(b)
	fmt.Println(str2)
}

// 类型断言 interface{}
// 倘若类型断言失败，则会抛出 panic，使用的时候，请千万注意处理。若不想让其抛出 panic，可以使用另一种断言语法
// 类型断言并不能用于两个通用类型的相互转换，只能用于将静态类型为 interface{} 类型的对象，转为具体的类型。
func AssertInterface() {
	//未知字段类型通过interface{}定义
	obj := make(map[string]interface{})
	var str string = "{\"num\":111,\"name\":\"jay\"}"
	_ = json.Unmarshal([]byte(str), &obj)
	//fmt.Println(obj["num"] + 1)
	//类型断言
	fmt.Println(obj["num"].(float64) + 1)
}
