package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

//select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
// 如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。

func main() {
	ch := make(chan string)
	go process(ch)

	//select {}  select 语句没有任何 case，它会一直阻塞，导致死锁。会触发 panic,(select 不用break;)

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	for {
		//无限循环
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
