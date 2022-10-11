package timer

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	//for range 循环从信道 ch 接收数据，直到该信道关闭。一旦关闭了 ch，for - range 循环会自动退出循环
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func OnlyChan() {
	ch := make(chan int)
	// 声明一个只能写入数据的通道类型, 并赋值为ch（指向chan关键词，为只写入通道，非指向chan的为只读通道）
	// 只有向左 <-符合 ，没有向右（->）符号，在chan左右边
	var chSendOnly chan<- int = ch
	//声明一个只能读取数据的通道类型, 并赋值为ch
	var chRecvOnly <-chan int = ch
	fmt.Println(chSendOnly, chRecvOnly)

	//make创建
	ch2 := make(<-chan int)
	fmt.Println(ch2)

	ch3 := make(chan<- int)
	fmt.Println(ch3)

	a := make(chan int)
	fmt.Println(a)

}

// 简单点说吧， 如果你需要写一个函数给别人调用， 假定这个功能就是源源不断的产生随机数。这样你可以先定义个双向通道，然后返回时候把他改成单项的，
// 这样外部调用者，就不能往chan，写入数据，减小出错的机率
