package main

//导入包。文件夹
import "oop/employee"

func main() {
	//使用包下面的New方法，小写没有导出，使用New返回对象（结构体），类似于Java构造方法
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
