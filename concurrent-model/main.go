package main

//@HanFa
import (
	"base_project/concurrent-model/runner"
	"fmt"
	"time"
)

//并发模型。规定时间内全部任务跑完
func createTask() func(int) {
	//返回实际执行的任务函数
	return func(i int) {
		time.Sleep(2 * time.Second)
		fmt.Printf("task: %d \n", i)
	}
}

func main() {
	runnerManager := runner.New(10 * time.Second)
	runnerManager.AddTasks(
		createTask(), createTask(), createTask(), createTask(), createTask(),
		createTask(), createTask(), createTask(), createTask(), createTask(),
		createTask(), createTask(), createTask(), createTask(),
	)
	err := runnerManager.Start()

	switch err {
	case runner.ErrInterrupt:
		fmt.Println("interrupt \n")
	case runner.ErrTimeOut:
		fmt.Println("timeout\n")
	default:
		fmt.Println("all task complete\n")
	}
}
