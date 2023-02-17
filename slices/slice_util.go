package slices

import (
	"sort"
	"strconv"
)

// 判断字符串是否存在切片
func InSliceString(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 判断数字是否存在切片
func InSliceInt(slice []int, item int) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 切片截取，从下标index开始，截取长度length; (slice[start_index:end_index]  原生截取，下标从0开始，截取包括start_index，不包括end_index)
func Cut(slice []string, startIndex int, length int) []string {
	//len()函数求长度，cap()函数求切片的容量
	maxLength := len(slice)
	endIndex := startIndex + length
	if endIndex > maxLength {
		endIndex = maxLength
	}
	return slice[startIndex:endIndex]
}

// 字符串切片合并
func MergeString(slice []string, slice2 []string) []string {
	//result := make([]string, len(slice)+len(slice2))
	//copy(result, slice)
	//copy(result[len(slice):], slice2)
	//...是解构数组
	return append(slice, slice2...)
}

// 整型切片合并
func MergeInt(slice []int, slice2 []int) []int {
	return append(slice, slice2...)
}

// 删除切片元素
func Unset(src []string, index int) []string {
	return append(src[:index], src[index+1:]...)
}

// 插入切片元素，向index位置插入一个item
func Insert(src []string, index int, item string) []string {
	srcNew := make([]string, len(src))
	copy(srcNew, src)
	return append(append(src[:index], item), srcNew[index:]...)
}

// 升序
func Sort(src []string) []string {
	sort.Strings(src)
	return src
}

// 倒序
func ReverseSort(src []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(src)))
	return src
}

// 获取两切片的交集
func Intersect(slice []string, slice2 []string) []string {
	m := make(map[string]int)
	result := make([]string, 0)
	for key, value := range slice2 {
		m[value] = key
	}
	for _, v := range slice {
		//使用map保存值判断
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// 切片去重
func Unique(src []string) []string {
	result := make([]string, 0, len(src))
	temp := map[string]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

//字符串转整型
func StringToInt(src []string) []int {
	var arr = make([]int, 0, len(src))
	if len(src) == 0 {
		return arr
	}
	for i, _ := range src {
		var num, _ = strconv.Atoi(src[i])
		arr = append(arr, num)
	}
	return arr
}

//整型转字符串
func IntToString(src []int) []string {
	var arr = make([]string, 0, len(src))
	if len(src) == 0 {
		return arr
	}
	for i, _ := range src {
		arr = append(arr, strconv.Itoa(src[i]))
	}
	return arr
}
