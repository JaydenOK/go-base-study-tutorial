package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

//并发写软件方案
//创建一个 channel 用来读和写这个随机数。
//创建 100 个生产者 goroutine。每个 goroutine 将产生随机数并将随机数写入到 channel 里。
//创建一个消费者 goroutine 用来从 channel 读取随机数并将它写入文件。这样的话我们就只有一个 goroutinue 向文件中写数据，从而避免竞争条件。
//一旦完成则关闭文件。

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			//有异常，退出
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	//通知写入完成
	done <- true
}

func main() {
	//信道保存要写入的数据
	data := make(chan int)
	//完成标记信道
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//多个生成数据协程
		go produce(data, &wg)
	}
	//一个消费写入文件协程
	go consume(data, done)
	go func() {
		//等待全部数据生成完,关闭data信道
		wg.Wait()
		//（关闭data信道后，go consume(data, done) 的 for d := range data 循环才会退出）
		close(data)
	}()

	//阻塞main函数，等待done值，等待所有写数据写完
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}

}
