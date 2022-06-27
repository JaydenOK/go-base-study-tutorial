package tests

import (
	"fmt"
	"testing"
)

//go选择排序法
func TestSort(t *testing.T) {
	var arr []int
	arr2 := append(arr, 5, 4, 3, 657, 76, 34, 845, 63, 5, 5, 657, 23, 24266, 2354)
	length := len(arr2)
	//最后一个不用比较
	for i := 0; i < length-1; i++ {
		//记录第一个值
		z := i
		j := i + 1
		for ; j < length; j++ {
			//遍历后面的，(比较待标记的值z)，是否还有比它小的数
			if arr2[j] < arr2[z] {
				z = j
			}
		}
		//即找到有比当前比较小的数
		if z != i {
			temp := arr2[z]
			arr2[z] = arr2[i]
			arr2[i] = temp
		}
	}
	fmt.Println(arr2)
}
