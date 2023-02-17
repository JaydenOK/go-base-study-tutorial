package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := map[string]string{"a": "11", "b": "22", "c": "33", "d": "666"}
	//m2 := map[string]string{"aa": "11", "bb": "22", "cc": "333", "a": "asadfaf"}
	//result := Merge(m1, m2)
	delete(m1, "aaa")
	fmt.Println(m1)
	//fmt.Println(result)
}

func demo() {
	//Map 是一种无序的键值对的集合。
	//Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	//Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的。
	//在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。
	//Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量

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
	// m := make(map[string]int)	// map 的正确声明，分配了实际的内存
	m["one"] = 1 // error: panic: assignment to entry in nil map
	//
	//golang中map是引用类型，应用类型的变量未初始化时默认的zero value是nil。直接向nil map写入键值数据会导致运行时错误panic: assignment to entry in nil map
	//
	//因为在声明dataMap后并未初始化它，所以它的值是nil, 不指向任何内存地址。需要通过make方法分配确定的内存地址。程序修改后即可正常运行:

	// slice 正确示例
	var s []int
	s = append(s, 1)
}

// 多维map的声明与实现方法
func multiMap() {
	//方法1 初始化一个空的多维映射
	mainMapA := map[string]map[string]string{}
	subMapA := map[string]string{"A_Key_1": "A_SubValue_1", "A_Key_2": "A_SubValue_2"}
	mainMapA["MapA"] = subMapA
	fmt.Println("MultityMapA")
	for keyA, valA := range mainMapA {
		for subKeyA, subValA := range valA {
			fmt.Printf("mapName=%s	Key=%s	Value=%s", keyA, subKeyA, subValA)
		}
	}

	//方法2 使用make声明一个多维映射(等同一般声明)
	//var mainMap map[string]map[string]string
	mainMapB := make(map[string]map[string]string)
	//内部容器必须再次初始化才能使用
	subMapB := make(map[string]string)
	subMapB["B_Key_1"] = "B_SubValue_1"
	subMapB["B_Key_2"] = "B_SubValue_2"
	mainMapB["MapB"] = subMapB
	fmt.Println("MultityMapB")

	for keyB, valB := range mainMapB {
		for subKeyB, subValB := range valB {
			fmt.Printf("mapName=%s	Key=%s	Value=%s", keyB, subKeyB, subValB)
		}
	}

	/* 方法3 使用interface{}初始化一个一维映射
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	//mainMapC := make(map[string]interface{})
	mainMapC := map[string]interface{}{}
	subMapC := make(map[string]string)
	subMapC["C_Key_1"] = "C_SubValue_1"
	subMapC["C_Key_2"] = "C_SubValue_2"
	mainMapC["MapC"] = subMapC
	fmt.Println("MultityMapC")
	for keyC, valC := range mainMapC {
		//此处必须实例化接口类型，即*.(map[string]string)
		//subMap := valC.(map[string]string)
		for subKeyC, subValC := range valC.(map[string]string) {
			fmt.Printf("mapName=%s	Key=%s	Value=%s", keyC, subKeyC, subValC)
		}
	}

}
