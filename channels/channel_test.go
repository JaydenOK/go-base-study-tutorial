package timer

import (
	"fmt"
	"testing"
	"time"
)

//Go语言的定时器实质是单向通道，time.Timer结构体类型中有一个time.Time类型的单向chan

func Test_timer(t *testing.T) {
	//Timer()
	//Tick()

	fmt.Println(uint64(time.Now().Unix()))

}

// 1，用NewTimer , Reset 创建的定时器 （go常用）
func Timer() {
	t1 := time.NewTimer(time.Second * 3) //从创建成功就开始计时了，3秒后，系统内部向只读通道 t1.C写入数据
	defer t1.Stop()
	for {
		//阻塞：等待3秒后，读取系统写入的数据
		<-t1.C
		func() {
			//要定时执行的程序代码块↓
			fmt.Println(NowTime())
		}()
		// 需要重置Reset 使 t1 重新开始计时
		t1.Reset(time.Second * 3)
	}
}

// 2，用time.Tick创建的定时器
func Tick() {
	//time.Tick() 函数返回的是一个 channel ，每隔指定的时间就会有数据从 channel 中出来。
	//range 其实是可以不接收遍历的值的，如果不接收，就可以简写成上面的格式。
	for range time.Tick(2 * time.Second) {
		func() {
			//要定时执行的程序代码块↓
			fmt.Println(NowTime())
		}()
	}
}

// 返回当前日期时间
func NowTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func testDemo() {
	/*
	   1.func NewTimer(d Duration) *Timer
	   创建一个计时器：d时间以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	*/
	//新建一个计时器：timer
	fmt.Println("1:", time.Now())
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("2:", time.Now())
	fmt.Printf("timer: %T\n", timer) //*time.Timer
	fmt.Println("3:", time.Now())
	//timer.C只读信道： 此处在等待channel中的信号，执行此段代码时会阻塞3秒
	ch2 := timer.C // <-chan time.Time
	fmt.Println("4:", time.Now())
	fmt.Println(<-ch2) //2022-05-17 00:08:50.503309 +0800 CST m=+3.000108752
	fmt.Println("5:", time.Now())
}
