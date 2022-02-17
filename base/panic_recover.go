package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func recovery() {
	//recover()恢复 panic
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
		//恢复 panic 之后，我们就失去了堆栈跟踪。 使用 Debug 包中的 PrintStack 函数。打印出堆栈跟踪
		debug.PrintStack()
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	go b() //panic 并不会恢复。因为调用 recovery 的协程和 b() 中发生 panic 的协程并不相同，因此不可能恢复 panic。panic仍然会发生
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
