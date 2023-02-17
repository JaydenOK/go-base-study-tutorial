package slices

import (
	"fmt"
	"sort"
	"testing"
)

type vvv []interface{}

func TestSlice(t *testing.T) {
	//创建切片
	//c := []int{6, 7, 8}
	a := []string{"a0", "b1", "a3", "c2", "d3", "e4", "a3", "d3"}
	//b := []string{"11", "22", "33", "44", "c2", "a0"}
	//c := []int{11, 22, 33}
	//d := []int{44, 55, 66, 77}
	//result := Insert(a, 4,"111")
	//slices := sort.StringSlice(a)
	fmt.Printf("%v", a)
	fmt.Println(sort.StringSlice(a))
	fmt.Println(ReverseSort(a))
}

func base() {
	//var names []string   //定义切片
	//var a [2]int  //定义数组

	//创建切片
	c := []int{6, 7, 8}
	fmt.Println(c)

	//从数组中创建切片: 使用语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片 (包左不包右)
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)       //[77 78 79]

	all := a[:] //切片获取所有数组元素
	fmt.Println(all)

	//切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中
	dArr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	dSlice := dArr[2:5]
	fmt.Println("array before", dArr)
	for i := range dSlice { //[2, 3, 4]
		dSlice[i]++ //改变的是原数组dArr下标为i的值
		fmt.Printf("i=%d,dSlice=%T \n", i, dSlice)
	}
	fmt.Println("array after", dArr)

	//切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数(被切原数组第一个索引起始)
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6

	//make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度
	s := make([]string, 5, 5)
	fmt.Println("s:", s)

	// func append（s[]T，x ... T）[]T  追加切片元素
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))             // capacity of cars is 3
	newCars := append(cars, "Toyota", "aaa")                                                       //追加，容量翻倍
	fmt.Println("newCars:", newCars, "has new length", len(newCars), "and capacity", cap(newCars)) // capacity of cars is doubled to 6

	// 使用 ... 运算符将一个切片添加到另一个切片
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	//切片包含长度、容量和指向数组第零个元素的指针。当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。
	// 因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见。
	//type slice struct {
	//	Length        int
	//	Capacity      int
	//	ZerothElement *byte
	//}

}

// 多维切片
func muitiSlice() {
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	//java : for in
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

// 函数传入切片，相当于传入引用类型
// 切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2] //neededCountries切片导致大数组不能回收
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //func copy(dst，src[]T)int 来生成一个切片的副本dst
	return countriesCpy
}

func study() {
	//	切片
	//	切片(Slice)是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。非常灵活，支持自动扩容。
	//
	//	切片属于引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	//
	//	切片的定义
	//	声明切片的基本语法：
	//
	//	var name []T
	//	其中：
	//
	//	name表示变量名
	//	T表示切片中的元素类型
	//	func main() {
	//		// 声明切片类型
	//		var a []string              //声明一个字符串切片
	//		var b = []int{}             //声明一个整型切片并初始化
	//		var c = []bool{false, true} //声明一个布尔切片并初始化
	//		var d = []bool{false, true} //声明一个布尔切片并初始化
	//		fmt.Println(a)              //[]
	//		fmt.Println(b)              //[]
	//		fmt.Println(c)              //[false true]
	//		fmt.Println(a == nil)       //true
	//		fmt.Println(b == nil)       //false
	//		fmt.Println(c == nil)       //false
	//		//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	//	}
	//	切片的长度和容量
	//	切片拥有自己的长度和容量，我们可以通过内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	//
	//	基于数组定义切片
	//	由于切片的底层就是一个数组，所以我们可以基于数组定义切片。
	//
	//	func main() {
	//		a := [5]int{55, 56, 57, 58, 59}
	//		b := a[1:4]
	//		fmt.Println(b)
	//		fmt.Printf("type of b:%T\n", b)
	//	}
	//	还支持如下方式：（与python类比）
	//
	//	c := a[1:] //[56 57 58 59]
	//	d := a[:4] //[55 56 57 58]
	//	e := a[:]  //[55 56 57 58 59]
	//	切片再切片
	//	func main() {
	//		//切片再切片
	//		a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	//		fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	//		b := a[1:3]
	//		fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	//		c := b[1:5]
	//		fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
	//	}
	//	输出：
	//
	//a:[北京 上海 广州 深圳 成都 重庆] type:[6]string len:6  cap:6
	//	b:[上海 广州] type:[]string len:2  cap:5
	//	c:[广州 深圳 成都 重庆] type:[]string len:4  cap:4
	//	注意： 对切片进行再切片时，索引不能超过原数组的长度，否则会出现索引越界的错误。
	//
	//	使用make()函数构造切片
	//	以上都是基于数组来创建切片的，如果要动态的创建一个切片，就需要使用make函数：
	//
	//	make([]T,size,cap)
	//
	//	其中：
	//
	//	T：切片的元素类型
	//	size：切片中元素的数量
	//	cap：切片的容量
	//	例子：
	//
	//	func main() {
	//	a := make([]int, 2, 10)
	//	fmt.Println(a)      //[0 0]
	//	fmt.Println(len(a)) //2
	//	fmt.Println(cap(a)) //10
	//	}
	//	上面代码中a的内部储存空间已经分配了10个，但是实际上值用了2个。容量并不影响当前元素的个数，所以len(a)返回2，cap(a)则返回该切片的容量。
	//
	//	切片的本质
	//	切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）、切片的容量（cap）。
	//
	//	举个例子，现在有一个数组a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}，切片s1 := a[:5]，相应示意图如下。
	//
	//
	//	image
	//	切片s2 := a[3:6]，相应示意图如下：
	//
	//
	//	image
	//	切片不能直接比较
	//	切片之间不能比较的，不能使用==操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较是和nil比较。一个nil值的切片并没有底层数组，
	//	一个nil值的切片的长度和容量都为0，但是不能说一个长度和容量都是0的切片一定是nil。
	//
	//	示例：
	//
	//	var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	//	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	//	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	//	要判断一个切片是否为空，要是用len(s) == 0来判断，不应该使用s == nil来判断。
	//
	//	切片的赋值拷贝
	//	下面代码演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容。
	//
	//	func main() {
	//	s1 := make([]int, 3) //[0 0 0]
	//	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	//	s2[0] = 100
	//	fmt.Println(s1) //[100 0 0]
	//	fmt.Println(s2) //[100 0 0]
	//	}
	//	切片遍历
	//	切片遍历和数组遍历是一致的，支持索引遍历和 for range遍历。
	//
	//	func main() {
	//	s := []int{1, 3, 5}
	//
	//	for i := 0; i < len(s); i++ {
	//	fmt.Println(i, s[i])
	//	}
	//
	//	for index, value := range s {
	//	fmt.Println(index, value)
	//	}
	//	}
	//	append()方法为切片添加元素
	//	Go语言的内建函数append()可以为切片动态添加元素。 每个切片会指向一个底层数组，这个数组能容纳一定数量的元素。当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在append()函数调用时。 举个例子：
	//
	//	func main() {
	//	//append()添加元素和切片扩容
	//	var numSlice []int
	//	for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	//	}
	//	}
	//	输出：
	//
	//	[0]  len:1  cap:1  ptr:0xc0000a8000
	//	[0 1]  len:2  cap:2  ptr:0xc0000a8040
	//	[0 1 2]  len:3  cap:4  ptr:0xc0000b2020
	//	[0 1 2 3]  len:4  cap:4  ptr:0xc0000b2020
	//	[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000b6000
	//	[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000b6000
	//	[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000b6000
	//	[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000b6000
	//	[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
	//	[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000
	//	从上面结果可以看出：
	//
	//	append()函数将元素追加到切片的最后并返回该切片
	//	切片的容量按照1,2,4,8,16的规则进行扩充，每次扩充后都是之前的2倍
	//	append()函数还支持一次性追加多个元素。 例如：
	//
	//	var citySlice []string
	//	// 追加一个元素
	//	citySlice = append(citySlice, "北京")
	//	// 追加多个元素
	//	citySlice = append(citySlice, "上海", "广州", "深圳")
	//	// 追加切片
	//	a := []string{"成都", "重庆"}
	//	citySlice = append(citySlice, a...)
	//	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
	//	切片的扩容策略
	//	可以通过查看$GOROOT/src/runtime/slice.go源码，其中扩容相关代码如下：
	//
	//	newcap := old.cap
	//	doublecap := newcap + newcap
	//	if cap > doublecap {
	//	newcap = cap
	//	} else {
	//	if old.len < 1024 {
	//	newcap = doublecap
	//	} else {
	//	// Check 0 < newcap to detect overflow
	//	// and prevent an infinite loop.
	//	for 0 < newcap && newcap < cap {
	//	newcap += newcap / 4
	//	}
	//	// Set newcap to the requested cap when
	//	// the newcap calculation overflowed.
	//	if newcap <= 0 {
	//	newcap = cap
	//	}
	//	}
	//	}
	//	由上所知：
	//
	//	首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
	//	否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
	//	否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
	//	如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	//	需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如int和string类型的处理方式就不一样。
	//
	//	使用copy()函数复制切片
	//	首先我们来看一个问题：
	//
	//	func main() {
	//	a := []int{1, 2, 3, 4, 5}
	//	b := a
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(b) //[1 2 3 4 5]
	//	b[0] = 1000
	//	fmt.Println(a) //[1000 2 3 4 5]
	//	fmt.Println(b) //[1000 2 3 4 5]
	//	}
	//	由于切片是引用类型，所以ａ和ｂ其实都指向了同一块内存地址。修改ｂ的同时ａ的值也会发生变化
	//
	//	Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
	//
	//	copy(destSlice, srcSlice []T)
	//	其中：
	//
	//	srcSlice: 数据来源切片
	//	destSlice: 目标切片
	//	例子：
	//
	//	func main() {
	//	// copy()复制切片
	//	a := []int{1, 2, 3, 4, 5}
	//	c := make([]int, 5, 5)
	//	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(c) //[1 2 3 4 5]
	//	c[0] = 1000
	//	fmt.Println(a) //[1 2 3 4 5]
	//	fmt.Println(c) //[1000 2 3 4 5]
	//	}
	//	从切片中删除元素
	//	Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：
	//
	//	func main() {
	//	// 从切片中删除元素
	//	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	//	// 要删除索引为2的元素
	//	a = append(a[:2], a[3:]...)
	//	fmt.Println(a) //[30 31 33 34 35 36 37]
	//	}
	//	总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
}
