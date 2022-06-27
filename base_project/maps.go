package main

import (
	"fmt"
	"reflect"
)

func main() {
	//make(map[type of key]type of value) 是创建 map 的语法。
	//personSalary := make(map[string]int)  //定义
	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	newEmp := "joe"
	// 获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
	value, ok := personSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	//检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可。
	//x := map[string]string{"one": "2", "two": "", "three": "3"}
	//if v := x["two"]; v == "" {
	//	fmt.Println("key two is no entry") // 键 two 存不存在都会返回的空字符串
	//}

	// for range 循环，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。
	personSalary2 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	fmt.Println("All items of a map")
	for key, value := range personSalary2 {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	//删除，函数没有返回值
	delete(personSalary2, "steve")

	//长度
	fmt.Println("length is", len(personSalary2))

	//和 slices 类似，map 也是引用类型,当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。

	//可以通过 reflect.DeepEqual 比较两个 slice/struct/map 是否相等:
	if reflect.DeepEqual(personSalary, personSalary2) {
		fmt.Println(personSalary, "==", personSalary2)
	}

	// map 错误示例
	var m map[string]int
	m["one"] = 1 // error: panic: assignment to entry in nil map
	// m := make(map[string]int)// map 的正确声明，分配了实际的内存

	// slice 正确示例
	var s []int
	s = append(s, 1)
}
