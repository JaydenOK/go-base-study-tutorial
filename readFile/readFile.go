package main

import (
	"fmt"
	"io/ioutil"
)

//IO util
func main() {
	path := "D:/www/goapp/src/readFile/log.txt"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
