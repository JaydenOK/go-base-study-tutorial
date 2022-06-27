package main

//因此在使用 GOPATH 模式下，我们需要将应用代码存放在固定的$GOPATH/src目录下，并且如果执行go get来拉取外部依赖会自动下载并安装到$GOPATH目录下。

import (
	"fmt"
	"reflect"
)

func main1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")
}
