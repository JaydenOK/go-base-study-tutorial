package main

import (
	"fmt"
	"sync"
)

//解决多个协程并发，修改数据错误问题  mutex:互斥的
////解决：协程函数传入访问锁（*sync.Mutex）

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1 //将竟态代码放在Lock 与 Unlock之间
	m.Unlock()
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

/*
也可以使用指定数量缓冲信道，限制协程数

package main
import (
    "fmt"
    "sync"
    )
var x  = 0
func increment(wg *sync.WaitGroup, ch chan bool) {
    ch <- true
    x = x + 1
    <- ch
    wg.Done()
}
func main() {
    var w sync.WaitGroup
    ch := make(chan bool, 1)  // ch := make(chan bool, 1)，只允许一个协程写入，其它协程
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go increment(&w, ch)
    }
    w.Wait()
    fmt.Println("final value of x", x)
}
*/
