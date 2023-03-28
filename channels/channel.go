//channel 及信道或管道
//无缓冲信道的发送和接收过程是阻塞的。（编程：一般是先开一个协程启动接收，后面启动另一个协程发送数据）
//有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。

//channel
//Go语言中的通道（channel）是一种特殊的类型。
//在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。
//
//通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
//
//（1）channel本身是一个队列，先进先出
//（2）线程安全，不需要加锁
//（3）本身是有类型的，string, int 等，如果要存多种类型，则定义成 interface类型
//（4）channel是引用类型，必须make之后才能使用，一旦 make，它的容量就确定了，不会动态增加！！它和map，slice不一样
//
//特点：
//（1）一旦初始化容量，就不会改变了。
//（2）当写满时，不可以写，取空时，不可以取。
//（3）发送将持续阻塞直到数据被接收
//把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示
//（4）接收将持续阻塞直到发送方发送数据。
//如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。
//（5）每次接收一个元素。
//通道一次只能接收一个数据元素。
//@todo 使用chan有限通道，阻塞投递，达到限制并发目的，控制任务并发数

package timer

import (
	"fmt"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println(squares, "+", cubes, "=", squares+cubes)
}

//在第 7 行，函数 calcSquares 计算一个数每位的平方和，并把结果发送给信道 squareop。与此类似，在第 17 行函数 calcCubes 计算一个数每位的立方和， 并把结果发送给信道 cubop。
//这两个函数分别在单独的协程里运行（第 31 行和第 32 行），每个函数都有传递信道的参数，以便写入数据。Go 主协程会在第 33 行等待两个信道传来的数据。
// 一旦从两个信道接收完数据，数据就会存储在变量 squares 和 cubes 里，然后计算并打印出最后结果。该程序会输出：
//Final output 1536

//这就需要用到信道转换（Channel Conversion）了。把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。
//package main
//import "fmt"
//func sendData(sendch chan<- int) {
//	sendch <- 10
//}
//func main() {
//	cha1 := make(chan int)
//	go sendData(cha1)
//	fmt.Println(<-cha1)
//}

//ch1 := make(chan int)                 // 创建一个整型类型的通道
//ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
//type Equip struct{ /* 一些字段 */ }
//ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip

//chan在使用make初始化时可附带一个可选参数来设置缓冲区。默认无缓冲，题目中便初始化的是无缓冲区的chan，这样只有写入的元素直到被读取后才能继续写入，不然就一直阻塞。
//ch := make(chan interface{}) 和 ch := make(chan interface{},1)是不一样的
//无缓冲的 不仅仅是只能向 ch 通道放 一个值 而是一直要有人接收，那么ch <- elem才会继续下去，要不然就一直阻塞着，也就是说有接收者才去放，没有接收者就阻塞。
//而缓冲为1则即使没有接收者也不会阻塞，因为缓冲大小是1只有当 放第二个值的时候 第一个还没被人拿走，这时候才会阻塞
