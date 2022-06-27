package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"unsafe"
)

//	fmt.Println(reflect.TypeOf(v))  打印变量v的类型
//GO字符串、数字转换：需引入"strconv"包
func stringConvertInt() {
	var s string = "123"
	//string到int   :  atoi (表示 ascii to integer)是把字符串转换成整型数
	i, err := strconv.Atoi(s)
	fmt.Println(i, err)

	//int到string
	s = strconv.Itoa(i)
	fmt.Println(s)

	//string到int64
	var str2 string = "10000000000000"
	i64, err := strconv.ParseInt(str2, 10, 64)
	fmt.Println(i64, err)

	//int64到string
	str2 = strconv.FormatInt(i64, 10)
	fmt.Println(str2)

	var b int
	fmt.Println(b)

	v := 1243
	fmt.Println(v)
}

func main() {
	stringConvertInt()
	return
	// var name type 是声明单个变量的语法。
	var a int = 12
	var b, c int = 100, 50 // 声明多个变量

	//批量声明
	var (
		name   = "naveen"
		age    = 29
		height int
	)

	//var cc int64 ;

	//name, age := "naveen", 29 // 简短声明,  :=  相当于php定义即用，= 需要之前有定义变量

	// b, c = 80, 90 // 给已经声明的变量b和c赋新值

	fmt.Println("num", a, b, c)

	fmt.Println("params:", name, age, height)

	// 格式说明符 %T 用于打印类型，而 %d 用于打印字节大小

	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
}

// Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换,可以手动转换
func testConvert() {
	// i := 55      //int
	// j := 67.8    //float64
	// sum := i + j //不允许 int + float64
	// sum := i + int(j) //j is converted to int 手动转换
}

func testConst() {
	const a = 55 // 允许
	// a = 89       // 不允许重新赋值

	//常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。
	// const b = math.Sqrt(4) // 不允许

}

func testConst2() {
	// 数字常量可以在表达式中自由混合和匹配，只有当它们被分配给变量或者在需要类型的代码中的任何地方使用时，才需要类型
	var a = 5.9 / 8
	fmt.Printf("a's type %T value %v", a, a)
}

func recursive() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	//第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10（10次遍历很快完成）。
	for i := 0; i < 10; i++ {
		//发生协程，10个协程都得到外部退出循环的值最终的i=10;
		go func() {
			fmt.Println("i1: ", i)
			wg.Done()
		}()
	}
	//第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i2: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//引用类型
//引用类型包括指针，slice切片，map ，chan，interface。
//变量直接存放的就是一个内存地址值，这个地址值指向的空间存的才是值。所以修改其中一个，另外一个也会修改（同一个内存地址）。
//引用类型必须申请内存才可以使用，make()是给引用类型申请内存空间。
//
//var a = []int{1,2,3,4,5}
//b := a      //此时a，b都指向了内存中的[1 2 3 4 5]的地址
//b[1] = 10   //相当于修改同一个内存地址，所以a的值也会改变
//c := make([]int,5,5)    //切片的初始化
//copy(c,a)   //将切片acopy到c
//c[1] = 20   //copy是值类型，所以a不会改变
//fmt.Printf("a的值是%v，a的内存地址是%p\n",a,&a)   //a的值是[1 10 3 4 5]，a的内存地址是0xc42000a180
//fmt.Printf("b的值是%v，b的内存地址是%p\n",b,&b)   //b的值是[1 10 3 4 5]，b的内存地址是0xc42000a1a0
//fmt.Printf("c的值是%v，c的内存地址是%p\n",c,&c)   //c的值是[1 20 3 4 5]，c的内存地址是0xc42000a1c0
//d := &a     //将a的内存地址赋值给d，取值用*d
//a[1] = 11
//fmt.Printf("d的值是%v，d的内存地址是%p\n",*d,d)   //d的值是[1 11 3 4 5]，d的内存地址是0xc420084060
//fmt.Printf("a的值是%v，a的内存地址是%p\n",a,&a)   //a的值是[1 11 3 4 5]，a的内存地址是0xc420084060
