package a

import (
	"fmt"
)

func main() {
	// price, no := 90, 6 // 定义 price 和 no,默认类型为 int
	// totalPrice := calculateBill(price, no)
	// fmt.Println("Total price is", totalPrice) // 打印到控制台上

	s1, s2 := rectProps(2.0, 3.0)

	fmt.Printf("%g : %v", s1, s2)

}

//函数声明 ,返回值的类型则定义在之后的 returntype (返回值类型)处
// func functionname(parametername type) returntype {
//     // 函数体（具体实现的功能）
// }

// 如果有连续若干个参数，它们的类型一致，那么我们无须一一罗列，只需在最后一个参数后添加该类型
func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// 多返回值（接收必须也是多个变量） a,b = rectProps()
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// 命名返回值
// 从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
// 请注意, 函数中的 return 语句没有显式返回任何值。由于 area 和 perimeter 在函数声明中指定为返回值, 因此当遇到 return 语句时, 它们将自动从函数返回。
func rectPropsAA(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
