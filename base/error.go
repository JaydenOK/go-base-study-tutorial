package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("D:\\www\\test.cc\\log.txt")

	//断言类型	v, ok := i.(int) //断言获取的是（int类型）
	//断言之后 判断ok的值;  if ok {}
	//断言底层结构体类型（类似于try catch finally异常）
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")

	errors.New("aa")
}
