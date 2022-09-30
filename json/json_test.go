package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_json(t *testing.T) {
	structJson()
	fmt.Println("####################################")
	mapJson()
	fmt.Println("####################################")
	decodeToStruct()
	fmt.Println("####################################")
	decodeToMap()
}

// 通过结构体生成json
type Persons struct {
	//"-"是忽略的意思          `json:"实际输出json的名称"`
	Name    string `json:"-"`
	Hobby   string `json:"hobby"`
	Company string `json:"company"`
}

// 一般用于绑定数据，mysql字段绑定,form表单绑定
func structJson() {
	p := Persons{"5lmh.com", "女", "eBay"}
	// 编码json，编码json使用json.Marshal()函数可以对一组数据进行JSON格式的编码
	// Marshal  :  v.	结集; 收集; 安排; 控制人群; 组织; 维持秩序
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err ", err)
	}

	fmt.Println("类型：", reflect.TypeOf(b)) //切片

	fmt.Println(string(b)) //切片转字符串

	// 带缩进格式化输出JSON
	//b, err = json.MarshalIndent(p, "", "     ")
	b, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("json err ", err)
	}
	fmt.Println(string(b))
}

// 通过map值interface生成json  (较通用)
func mapJson() {
	student := make(map[string]interface{})
	student["name"] = "5lmh.com"
	student["age"] = 18
	student["sex"] = "man"
	b, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("类型：", reflect.TypeOf(b)) //切片
	fmt.Println(string(b))
}

func decodeToStruct() {
	var str string = `{"hobby":"aaa","company":"ccc"}`
	b := []byte(str)
	var p Persons
	err := json.Unmarshal(b, &p) //解析到P结构体
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("类型：", reflect.TypeOf(p))          //具体类型名称： main.Persons
	fmt.Println("类型2：", reflect.ValueOf(p).Kind()) //类型2： struct
	fmt.Println(p)
}

func decodeToMap() {
	// 假数据
	// int和float64都当float64
	b := []byte(`{"age":1.3,"name":"5lmh.com","marry":false}`)
	// 声明接口
	var i interface{}
	err := json.Unmarshal(b, &i) //解析json到接口i (map数据结果)
	if err != nil {
		fmt.Println(err)
	}
	// 自动转到map
	fmt.Println(i)
	// 可以判断类型
	m := i.(map[string]interface{})
	for k, v := range m {
		//v.(type)  取得类型
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		case bool:
			fmt.Println(k, "是bool类型", vv)
		default:
			fmt.Println("其他")
		}
	}
}
