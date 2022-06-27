package main

//当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止。
// panic("runtime error: last name cannot be nil") ; 类似于 throw exception panic()  finally ( print "runtime error: last name cannot be nil") 后面的语句不在执行
// defer 在发生 panic 的函数中依然能执行，即返回panic前，执行了defer

import (
	"fmt"
)

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

//This program prints,
//
//deferred call in fullName
//deferred call in main
//panic: runtime error: last name cannot be nil     (程序控制到达了顶层函数，因此该函数会打印出 panic 信息，然后是堆栈跟踪，最后终止程序。)
//
//goroutine 1 [running]:
//main.fullName(0x1042bf90, 0x0)
///tmp/sandbox060731990/main.go:13 +0x280
//main.main()
///tmp/sandbox060731990/main.go:22 +0xc0
