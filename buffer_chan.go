package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		//写2个后阻塞，等待被读取
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

//在上面的程序中，第 16 行在 Go 主协程中创建了容量为 2 的缓冲信道 ch，而第 17 行把 ch 传递给了 write 协程。接下来 Go 主协程休眠了两秒。在这期间，write 协程在并发地运行。write 协程有一个 for 循环，依次向信道 ch 写入 0～4。而缓冲信道的容量为 2，因此 write 协程里立即会向 ch 写入 0 和 1，接下来发生阻塞，直到 ch 内的值被读取。因此，该程序立即打印出下面两行：
//
//successfully wrote 0 to ch
//successfully wrote 1 to ch
//打印上面两行之后，write 协程中向 ch 的写入发生了阻塞，直到 ch 有值被读取到。而 Go 主协程休眠了两秒后，才开始读取该信道，因此在休眠期间程序不会打印任何结果。主协程结束休眠后，在第 19 行使用 for range 循环，开始读取信道 ch，打印出了读取到的值后又休眠两秒，这个循环一直到 ch 关闭才结束。所以该程序在两秒后会打印下面两行：
//
//read value 0 from ch
//successfully wrote 2 to ch
//该过程会一直进行，直到信道读取完所有的值，并在 write 协程中关闭信道。最终输出如下：
//
//successfully wrote 0 to ch
//successfully wrote 1 to ch
//read value 0 from ch
//successfully wrote 2 to ch
//read value 1 from ch
//successfully wrote 3 to ch
//read value 2 from ch
//successfully wrote 4 to ch
//read value 3 from ch
//read value 4 from ch

//
func test() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve" //第三个不能被写入，死锁
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
