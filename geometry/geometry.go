// geometry.go
package main

import (
	"fmt"
	//"geometry/rectangle" // 导入自定义包
	_ "geometry/rectangle" // 导入了，暂时不使用方法
	"log"
)

//有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。例如，我们或许需要确保调用了 rectangle 包的 init 函数，而不需要在代码中使用它。
// 这种情况也可以使用空白标识符 _ ，如下所示。

//main 包的初始化顺序为：
//
//首先初始化被导入的包。因此，首先初始化了 rectangle 包。
//接着初始化了包级别的变量 rectLen 和 rectWidth。
//调用 init 函数。
//最后调用 main 函数。

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	//fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	//fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
}
