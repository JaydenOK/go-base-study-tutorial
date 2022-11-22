package a

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
//go语言http
1、net/http 包提供的 http.ListenAndServe() 方法，在指定的地址监听
该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连 接请求。该方法有两个参数：第一个参数 addr 即监听地址；
第二个参数表示服务端处理程序， 通常为空，这意味着服务端调用 http.DefaultServeMux 进行处理，
而服务端编写的业务逻辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中。

2.处理https请求
func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error

3.路由
http.HandleFunc()方法接受两个参数
第一个参数是HTTP请求的 目标路径"/hello"，该参数值可以是字符串，也可以是字符串形式的正则表达式
第二个参数指定具体的回调方法，比如helloHandler。
当我们的程序运行起来后，访问http://localhost:8080/hello ， 程序就会去调用helloHandler()方法中的业务逻辑程序。
*/
//开启一个http服务
func httpServe() {
	//业务逻辑处理程序,访问：http://localhost:8080/list，http://localhost:8080/add 查看
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("page list"))
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(302)
		w.Write([]byte("page add"))
	})
	fmt.Println("启动服务")

	//监听并启动服务地址、端口
	var ip = "127.0.0.1"
	var port = 8080
	http.ListenAndServe(ip+":"+strconv.Itoa(port), nil)

}

func get() {
	//模拟一个get提交请求
	resp, err := http.Get("https://www.sina.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println("#################################################################")
}

func post() {
	//模拟一个post提交请求
	resp2, err2 := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err2 != nil {
		panic(err2)
	}
	//最后关闭连接
	defer resp2.Body.Close()
	//读取报文中所有内容
	body2, err2 := ioutil.ReadAll(resp2.Body)
	//输出内容
	fmt.Println(string(body2))
	fmt.Println("#################################################################")
}

func main() {
	//get()
	//post()
	httpServe()
}
