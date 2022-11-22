package a

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	//一定在返回时执行 wg.Done(),不用再每个 if 语句都写
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// 含有 defer 语句的函数，会在该函数将要返回之前，调用defer另一个函数
//在上述程序中的第 11 行，for range 循环会遍历一个字符串，并在第 12 行调用了 defer fmt.Printf("%c", v)。这些延迟调用会添加到一个栈中，
// 按照后进先出的顺序执行，因此，该字符串会逆序打印出来。该程序会输出：
//Orignal String: Naveen
//Reversed String: neevaN

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main22() {

	a := 1                               //line 1
	b := 2                               //2
	defer calc("1", a, calc("10", a, b)) //3
	a = 0                                //4
	defer calc("2", a, calc("20", a, b)) //5
	b = 1                                //6
	//答 输出结果为：
	//10 1 2 3
	//20 0 2 2
	//2 0 2 2
	//1 1 3 4

	//解析
	//在解题前需要明确两个概念：
	//defer是在函数末尾的return前执行，先进后执行，具体见问题1。
	//函数调用时 int 参数发生值拷贝。

	//不管代码顺序如何，defer calc func中参数b必须先计算，故会在运行到第三行时，执行calc("10",a,b)输出：10 1 2 3得到值3，将cal("1",1,3)存放到延后执执行函数队列中。

	//执行到第五行时，现行计算calc("20", a, b)即calc("20", 0, 2)输出：20 0 2 2得到值2,将cal("2",0,2)存放到延后执行函数队列中。
	//执行到末尾行，按队列先进后出原则依次执行：cal("2",0,2)、cal("1",1,3) ，依次输出：2 0 2 2、1 1 3 4 。
}
