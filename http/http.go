//main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//file, err := os.Open("./hello.txt")
	//if err != nil {
	//	fmt.Println("Open file err = ", err)
	//}
	//defer file.Close()
	//const (
	//	defaultBufSize = 4096
	//)
	//reader := bufio.NewReader(file)
	//for  {
	//	str ,err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		break
	//	}
	//	_, _ = fmt.Fprintln(w,string(str) )
	//}
	b, _ := ioutil.ReadFile("./hello.html") //内容不是特别多
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http serve failed, err:%v\n", err)
		return
	}
}
