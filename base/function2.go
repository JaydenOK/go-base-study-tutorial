package main

import "fmt"

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)

	//在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片
	nums := []int{89, 90, 95}
	find(89, nums...)
}

/**
可变参数函数的工作原理是把可变参数转换为一个新的切片。以上面程序中的第 22 行为例，find 函数中的可变参数是 89，90，95 。
find 函数接受一个 int 类型的可变参数。因此这三个参数被编译器转换为一个 int 类型切片 int []int{89, 90, 95} 然后被传入 find函数。
*/
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	//相当于循环切片
	//foreach as key=>value
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
