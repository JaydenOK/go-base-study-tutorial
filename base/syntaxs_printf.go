package main

import "fmt"

/**
Println：(终端中有显示)
1. 用默认的类型格式T显示方式将传入的参数写入到标准输出里面
2. 多个传入参数之间使用空格分隔，
3. 在显示的最后追加换行符，
4. 返回值为 写入标准输出的字节数和写入过程中遇到的问题。

Printf：(终端中有显示)
1. 用传入的格式化规则符将传入的变量写入到标准输出里面
2. 返回值为 写入标准输出的字节数和写入过程中遇到的问题。

Sprintf: (终端中不会显示,格式化字符串使用)
1. 用传入的格式化规则符将传入的变量格式化，
2. 返回为 格式化后的字符串。

格式化规则符如下：
	%v 输出结构体 {10 30}
	%+v 输出结构体显示字段名 {one:10 tow:30}
	%#v 输出结构体源代码片段 main.Point{one:10, tow:30}
	%T 输出值的类型            main.Point
	%t 输出格式化布尔值      true
	%d`输出标准的十进制格式化 100
	%b`输出标准的二进制格式化 99 对应 1100011
	%c`输出定整数的对应字符  99 对应 c
	%x`输出十六进制编码  99 对应 63
	%f`输出十进制格式化  99 对应 63
	%e`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	%E`输出科学技科学记数法表示形式  123400000.0 对应 1.234000e+08
	%s 进行基本的字符串输出   "\"string\""  对应 "string"
	%q 源代码中那样带有双引号的输出   "\"string\""  对应 "\"string\""
	%p 输出一个指针的值   &jgt 对应 0xc00004a090
	% 后面使用数字来控制输出宽度 默认结果使用右对齐并且通过空格来填充空白部分
	%2.2f  指定浮点型的输出宽度 1.2 对应  1.20
	%*2.2f  指定浮点型的输出宽度对齐，使用 `-` 标志 1.2 对应  *1.20
*/

/*
GOROOT：Golang的安装路径。
GOPATH：可以理解为工作目录或者工作区，也是平时接触最多的一个变量。它可以是一个目录，可以是多个目录路径，每个目录代表一个工作区。
这些目录用于放置Go语言的源码文件（src），以及安装（命令go install）后的归档文件（pkg目录）和可执行文件（bin目录）。
GOBIN：GOROOT目录下的可执行文件放置目录，一般指bin。
GOPATH=D:\www\goapp 为工作目录;   D:\www\goapp\src 源码文件夹
Go 的 unsafe 包提供了一个 Sizeof 函数，该函数接收变量并返回它的字节大小
*/

func main() {
	//if
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	//for
	for i := 1; i <= 20; i++ {
		if i > 16 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	//多值语句
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//无限循环,使用break退出循环 ,没有while循环
	j := 0
	for {
		fmt.Printf("Hello World: %d \n", j)
		j++
		if j > 10 {
			break
		}
	}

	//switch
	var aa = 10
	switch aa {
	case 1, 2, 3, 4, 5: // 一个选项多个表达式
		fmt.Println("in 1-5")
	case 6, 7, 8, 9, 10:
		fmt.Println("in 6-10")
	default:
		fmt.Println("aa not in 1-10")
	}

	//switch 和 case 的表达式不一定是常量，使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

}

func number() int {
	num := 15 * 5
	return num
}
