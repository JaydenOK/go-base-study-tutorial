package pool

import (
	"fmt"
	"reflect"
	"sync"
)

// Pool implements worker pool
type Pool interface {
	// Delegate job to a workers
	// will block if channel is full, you might want to wrap it with goroutine to avoid it
	// will panic if called after Stop()
	Delegate(args ...interface{}) error

	// AddWorker adds worker to the pool
	AddWorker(fn interface{}) error
	// RemoveWorker removes worker from the pool
	RemoveWorker(fn interface{}) error

	// WorkersNum returns number of workers in the pool
	WorkersNum() int

	// Stop removes all workers workers
	// to resume work add them again
	Stop()
}

type quitCh chan struct{}
type workers map[reflect.Value][]quitCh

type pool struct {
	queue   chan []reflect.Value
	workers workers
	mtx     sync.RWMutex
}
//委托
func (p *pool) Delegate(args ...interface{}) error {
	if len(p.workers) == 0 {
		return fmt.Errorf("there is no workers in pool")
	}

	p.queue <- buildQueueValue(args)

	return nil
}

func (p *pool) AddWorker(fn interface{}) error {
	if err := isValidHandler(fn); err != nil {
		return err
	}

	//反射函数
	worker := reflect.ValueOf(fn)

	//1）多个读锁可以同时操作 .
	//2）写锁一旦加锁，不管是读操作还是写操作都不能被执行。
	//3）读锁加锁了之后，只要读锁还没有解锁，写操作是不能被执行的。

	//Go 中的 map 是不支持 并发写的，我们可以利用 读写锁 RWMutex 来实现并发安全的 map。在读多写少的情况下，使用 RWMutex 要比 Mutex 性能高。
	p.mtx.Lock()
	defer p.mtx.Unlock()

	q := make(quitCh)

	if _, ok := p.workers[worker]; !ok {
		p.workers[worker] = []quitCh{q}
	} else {
		p.workers[worker] = append(p.workers[worker], q)
	}

	//阻塞，等待队列通道p.queue投递数据
	go func() {
		for {
			select {
			case args := <-p.queue:
				//反射后，执行函数
				worker.Call(args)
			case <-q:
				return
			}
		}
	}()

	return nil
}

func (p *pool) RemoveWorker(fn interface{}) error {
	if err := isValidHandler(fn); err != nil {
		return err
	}

	worker := reflect.ValueOf(fn)

	p.mtx.Lock()
	defer p.mtx.Unlock()

	if len(p.workers[worker]) > 0 {
		close(p.workers[worker][len(p.workers[worker])-1])

		p.workers[worker] = p.workers[worker][:len(p.workers[worker])-1]
	} else {
		delete(p.workers, worker)
	}

	return nil
}

func (p *pool) WorkersNum() int {
	sum := 0
	for _, qChs := range p.workers {
		sum += len(qChs)
	}

	return sum
}

func (p *pool) Stop() {
	for _, qChs := range p.workers {
		for _, ch := range qChs {
			close(ch)
		}
	}
}

func isValidHandler(fn interface{}) error {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return fmt.Errorf("%s is not a reflect.Func", reflect.TypeOf(fn))
	}

	return nil
}

func buildQueueValue(args []interface{}) []reflect.Value {
	reflectedArgs := make([]reflect.Value, 0)

	for _, arg := range args {
		reflectedArgs = append(reflectedArgs, reflect.ValueOf(arg))
	}

	return reflectedArgs
}

// New creates new worker pool with a given job queue length
func New(queueLength int) Pool {
	return &pool{
		queue:   make(chan []reflect.Value, queueLength),
		workers: make(workers),
	}
}
