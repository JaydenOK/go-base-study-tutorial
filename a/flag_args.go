package a

import (
	"fmt"
	"os"
)

// os.Args demo
func main() {
	//os.Args是一个[]string
	//os.Args是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称。
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

//
//$ ./args_demo a b c d
//args[0]=./args_demo
//args[1]=a
//args[2]=b
//args[3]=c
//args[4]=d
