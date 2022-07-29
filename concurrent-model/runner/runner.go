package runner

//@HanFa
import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeOut   = errors.New("task running timeout")
	ErrInterrupt = errors.New("interrupt from os")
)

//任务管理模型
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int) //添加的任务，切片，多个，每个任务即是函数
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t),
		tasks:     make([]func(int), 0),
	}
}

//添加多个任务到任务管理模型
func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行遍历每个任务
func (r *Runner) run() error {
	//有限循环模式
	for id, task := range r.tasks {
		select {
		//如果有系统终止信号，则立即中断执行后面的任务，并返回
		case <-r.interrupt:
			signal.Stop(r.interrupt) //让 r.interrupt 停止接收其它信号
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}

//开启任务管理器里的任务，开始执行
func (r *Runner) Start() error {
	//1，监听系统终端信号，将系统信号传入runner管理器的信号上，r.interrupt缓存为1，只保存一个，添加多个，阻塞时会丢弃后面的信号
	signal.Notify(r.interrupt, os.Interrupt)
	//2，启动协程执行任务
	go func() {
		r.complete <- r.run()
	}()
	//3，等待执行结果，中断或者全部执行完
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}
