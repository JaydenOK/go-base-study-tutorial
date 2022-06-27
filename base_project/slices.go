package main

import "fmt"

func main() {

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

//多维切片
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

//函数传入切片，相当于传入引用类型
//切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2] //neededCountries切片导致大数组不能回收
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //func copy(dst，src[]T)int 来生成一个切片的副本dst
	return countriesCpy
}
