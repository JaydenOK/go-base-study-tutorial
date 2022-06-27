package main

import (
	"fmt"
	"unicode/utf8"
)

//Go 语言中的字符串是一个字节切片。可以像切片一样使用，下标访问修改，Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x-", s[i]) //	%x输出十六进制编码
		//fmt.Printf("%c-", s[i])  	// %c输出字符
	}
}

//rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
func printChars(s string) {
	//传入字符串，字符串切片 转成 rune切片
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c-", runes[i])
	}
}

func main() {
	name := "Hello World"
	printBytes(name) //48-65-6c-6c-6f-20-57-6f-72-6c-64-
	fmt.Printf("\n")

	printChars(name) //H-e-l-l-o- -W-o-r-l-d-
	fmt.Printf("\n\n")
	//这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码占用了 c3 和 b1 两个字节。它的 UTF-8 编码占用了两个字节 c3 和 b1。而我们打印字符时，
	//却假定每个字符的编码只会占用一个字节，这是错误的。在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间

	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")

	printChars(name)
	fmt.Printf("\n")

	//for range 循环字符串，循环返回的是是当前 rune 的字节位置
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

	runeToString()
	editString()
}

//十六进制字节切片 转 字符串，直接十进制转也可以
func toString() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

//rune双字节切片 转 字符串
func runeToString() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)                         //Señor
	fmt.Println(len(str))                    //6
	fmt.Println(utf8.RuneCountInString(str)) //5
}

//go 字符串 修改方法(Go 中的字符串是不可变的)
func editString() {
	h := "hello"
	//不能直接修改字符串  h[0] = 'a' (×)
	runeString := []rune(h)
	fmt.Println("runeSring:", runeString)
	runeString[0] = 'a'
	h = string(runeString)
	fmt.Println("h:", h)
}
