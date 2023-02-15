package arrays

import (
	"fmt"
	"strconv"
	"strings"
)

func InArray(search interface{}, arr [...]interface{}) bool {

	return false
}

func base() {
	//初始化默认[0,0],//数组的大小是类型的一部分。因此 [5]int 和 [25]int 是不同类型。数组不能调整大小，不要担心这个限制
	var a [2]int
	fmt.Println("初始化:", a)

	//简略声明，在简略声明中，不需要将数组中所有的元素赋值。多余的赋值为0
	b := [5]int{12, 78, 50}
	fmt.Println("简略声明:", b)

	//自动决定长度
	c := [...]int{12, 78, 50}
	fmt.Println("自动决定长度:", len(c), c)

	//for循环
	d := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(d); i++ { // looping from 0 to the length of the array
		fmt.Printf("index:%d for循环: %.2f\n", i, d[i])
	}

	// i 为固定的地址
	//for i, _ := range stus {
	//			m[stu.Name] = &i   (错误取法)

	//	stu:=stus[i]
	//	m[stu.Name] = &stu

	//}

	//range循环 类似 list()
	//下标 index可用空白标识符: for _, v := range a { }
	e := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for index, value := range e { //range returns both the index and value
		fmt.Printf("index: %d range循环: %.2f\n", index, value)
		sum += value
	}

	//数组
	var arr1 [5]string
	fmt.Println(arr1)

	//切片
	var arr2 []string
	fmt.Println(arr2)

}




//合并数组
func MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

//删除数组
func DeleteArray(src []interface{}, index int) (result []interface{}) {
	result = append(src[:index], src[(index+1):]...)
	return
}

// []string => []int
func ArrayStr2Int(data []string) []int {
	var (
		arr = make([]int, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		var num, _ = strconv.Atoi(data[i])
		arr = append(arr, num)
	}
	return arr
}

// []int => []string
func ArrayInt2Str(data []int) []string {
	var (
		arr = make([]string, 0, len(data))
	)
	if len(data) == 0 {
		return arr
	}
	for i, _ := range data {
		arr = append(arr, strconv.Itoa(data[i]))
	}
	return arr
}

// str[TrimSpace] in string list
func TrimSpaceStrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == strings.TrimSpace(row) {
				return true
			}
		}
	}
	return false
}

// str in string list
func StrInArray(str string, data []string) bool {
	if len(data) > 0 {
		for _, row := range data {
			if str == row {
				return true
			}
		}
	}
	return false
}

// str in int list
func IntInArray(num int, data []int) bool {
	if len(data) > 0 {
		for _, row := range data {
			if num == row {
				return true
			}
		}
	}
	return false
}