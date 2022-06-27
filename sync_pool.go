package main

/**
sync.Pool 是临时对象池，存储的是临时对象，不可以用它来存储 socket 长连接和数据库连接池等。
sync.Pool 本质是用来保存和复用临时对象，以减少内存分配，降低 GC 压力，比如需要使用一个对象，就去 Pool 里面拿，如果拿不到就分配一份，这比起不停生成新的对象，用完了再等待 GC 回收要高效的多。
*/
import (
	"fmt"
	"sync"
)

type syncStudent struct {
	Name string
	Age  int
}

var syncStudentPool = &sync.Pool{
	New: func() interface{} {
		return new(syncStudent)
	},
}

func syncNew(name string, age int) *syncStudent {
	stu := syncStudentPool.Get().(*syncStudent)
	stu.Name = name
	stu.Age = age
	return stu
}

func syncRelease(stu *syncStudent) {
	stu.Name = ""
	stu.Age = 0
	syncStudentPool.Put(stu)
}

func main() {
	//当使用 syncStudent 对象时，只需要调用 New() 方法获取对象，获取之后使用 defer 函数进行释放即可。
	stu := syncNew("tom", 30)
	defer syncRelease(stu)
	a := new(int)
	fmt.Println(a)
	//业务代码。。。

}
