package generic

import "testing"

//泛型测试 version > go1.18

// 使用泛型
//func GetMaxNum[T int | int8](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}

// 像声明接口一样声明
//type MyInt interface {
//	int | int8 | int16 | int32 | int64
//}
//
//// T的类型为声明的MyInt
//func GetMaxNum[T MyInt](a, b T) T {
//	if a > b {
//		return a
//	}
//	return b
//}
/////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////////////////////////////



func TestA(t *testing.T) {
	//var a int = 10
	//var b int = 20
	//
	//// 方法1，正常调用，编译器会自动推断出传入类型是int
	//GetMaxNum(a, b)
	//
	//// 方法2，显式告诉函数传入的类型是int
	//GetMaxNum[int](a, b)


	/////////////////////////////////////////////////////////////////////////////////////////////////////////

	// 声明要使用的泛型的类型
	var u User[int8, string]
	// 赋值
	u.age = 18
	u.name = "weiwei"
	// 调用方法
	age := u.GetAge()
	name := u.GetName()
	// 输出结果 18 weiwei
	fmt.Println(age, name)
}


//泛型与结构体
//创建一个带有泛型的结构体User，提供两个获取age和name的方法
//注意：只有在结构体上声明了泛型，结构体方法中才可以使用泛型
type AgeT interface {
	int8 | int16
}
type NameE interface {
	string
}
type User[T AgeT, E NameE] struct {
age  T
name E
}
// 获取age
func (u *User[T, E]) GetAge() T {
	return u.age
}


// 获取name
func (u *User[T, E]) GetName() E {
	return u.name
}