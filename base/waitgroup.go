package main

//工作池

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

//算法为，对number每个位数，进行各位相加（如:634=4+3+6 算法返回sum=13）
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second) //模拟延迟2秒返回
	return sum
}
func worker(wg *sync.WaitGroup) {
	//取出job，延迟2秒返回，输入到results信道  (此处取出job，allocate 协程，取消阻塞，继续移入数据)
	//输出信息的延迟效果主要是digits 方法 ，模拟了延迟效果
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done() //完成一个，协程池计数器-1
}

/**
@param noOfWorkers
开启worker协程数
*/
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup //1，创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)      // 计数器 +1
		go worker(&wg) //传递 wg 的地址是很重要的。如果没有传递 wg 的地址，那么每个 Go worker协程将会得到一个 WaitGroup 值的拷贝，因而当它们执行结束时，调用函数createWorkerPool并不会知道。
	}
	wg.Wait()      //Wait() 方法会阻塞调用它的 Go 协程，直到计数器变为 0 后才会停止阻塞。(等待所有worker协程执行完，go worker(&wg))
	close(results) //执行完关闭results信道, done返回结果
}

//初始化 chan
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job //阻塞，直到有数据移除信道后
	}
	close(jobs)
}

func result(done chan bool) {
	//for range 循环从信道 results 接收数据，直到该信道关闭。一旦关闭了 results，循环会自动结束
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

//全局变量
var jobs = make(chan Job, 10)       //缓冲信道10
var results = make(chan Result, 10) //缓冲信道10

//缓冲信道的重要应用之一就是实现工作池。
//一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。`

func main() {
	//context.WithCancel()
	startTime := time.Now()
	go allocate(100) //开启-将100个结构体存放入缓冲信道(jobs)协程  (不等待，下面继续执行，里边缓冲信道jobs阻塞)
	done := make(chan bool)
	go result(done) //不等待，for range (不等待，下面继续执行，里边缓冲信道results阻塞)
	//主消费流程
	createWorkerPool(10) //createWorkerPool 会关闭上面的协程(go result)
	<-done               //done 变为true, result协程才执行完。阻塞：main 函数会监听 done 信道的通知，等待所有结果打印结束。
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
