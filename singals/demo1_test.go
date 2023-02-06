package singals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

//前言
//信号(Signal)是Linux, 类Unix和其它POSIX兼容的操作系统中用来进程间通讯的一种方式。对于Linux系统来说，信号就是软中断，用来通知进程发生了异步事件。
//
//当信号发送到某个进程中时，操作系统会中断该进程的正常流程，并进入相应的信号处理函数执行操作，完成后再回到中断的地方继续执行。
//
//有时候我们想在Go程序中处理Signal信号，比如收到SIGTERM信号后优雅的关闭程序，以及 goroutine结束通知等。
//
//Go 语言提供了对信号处理的包（os/signal）。
//
//Go 中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。
//
//Go信号通知机制可以通过往一个channel中发送os.Signal实现。

func Test1(t *testing.T) {
	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//创建一个bool channel
	done := make(chan bool, 1)

	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	//程序将在此处等待，直到它预期信号（如Goroutine所示）
	//在“done”上发送一个值，然后退出。
	fmt.Println("awaiting signal")
	//阻塞主协程
	<-done
	fmt.Println("exiting")
}

func signalLists() {
	//信号	值	动作	说明
	//SIGHUP	1	Term	终端控制进程结束(终端连接断开)
	//SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
	//SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
	//SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
	//SIGABRT	6	Core	调用abort函数触发
	//SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
	//SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
	//SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
	//SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
	//SIGALRM	14	Term	时钟定时信号
	//SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
	//SIGUSR1	30,10,16	Term	用户保留
	//SIGUSR2	31,12,17	Term	用户保留
	//SIGCHLD	20,17,18	Ign	子进程结束(由父进程接收)
	//SIGCONT	19,18,25	Cont	继续执行已经停止的进程(不能被阻塞)
	//SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)
	//SIGTSTP	18,20,24	Stop	停止进程(可以被捕获、阻塞或忽略)
	//SIGTTIN	21,21,26	Stop	后台程序从终端中读取数据时触发
	//SIGTTOU	22,22,27	Stop	后台程序向终端中写数据时触发
}
