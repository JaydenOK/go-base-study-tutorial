package main

import (
	"fmt"
	"time"
)

//select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
//select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

func set(ch chan string) {
	time.Sleep(1)
	ch <- "s"
}

func setInt(ch chan int) {
	ch <- 5
}

func main() {
	//select 会循环检测条件，如果有满足则执行并退出，否则一直循环检测
	ch := make(chan string)
	ch2 := make(chan int)
	go set(ch)
	go setInt(ch2)
	select {
	case a := <-ch:
		fmt.Println("ch1:", a)
	case b := <-ch2:
		fmt.Println("ch2:", b)
	}

	//switch 不用加break
	b := 2
	switch a := b + 1; a {
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("xxx")
	}
}
