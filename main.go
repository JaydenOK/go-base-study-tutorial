package main

import (
	"base_project/package1"
	"fmt"
)

/**
使用go module模式
即使放在src目录下，src/project1,src/project2；每个目录下一个go.mod（命令:go mod init生成） 和 main.go文件
src/project1/package1为package1包，其下文件的包名即为package package1(包名与目录保持一致)
*/
func main() {
	fmt.Println("hi")
	package1.Import()
}
