package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

//按字节数读取 : 	b := make([]byte, 3)

func main() {
	//使用命令行标记来传递文件路径: bin/readFile2.exe -fpath=/path-of-file/test.txt
	fptr := flag.String("fpath", "log.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	//最后关闭文件流
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		_, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b))
	}
}

//按行读取
func main2() {
	fptr := flag.String("fpath", "log.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//按行扫描读取
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
