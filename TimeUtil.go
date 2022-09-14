package main

import (
	"fmt"
	"time"
)

func main() {
	nowDate := NowTime()
	fmt.Println(nowDate)

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板，神奇模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, nowDate, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	fmt.Println(timestamp)

}

func NowTime() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
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
