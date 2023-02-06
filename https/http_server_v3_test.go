package https

import (
	"log"
	"net/http"
	"testing"
	"time"
)

// v3 这里可以自定义http服务器配置，都在Server这个结构体中,这个对象能配置监听地址端口，配置读写超时时间，配置handler,配置请求头最大字节数...，所有稍微改造一下v2的程序得到v3版:
func TestV3(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler3{})
	mux.HandleFunc("/bye", sayBye3)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	log.Println("Starting v3 httpserver")
	log.Fatal(server.ListenAndServe())
}

type myHandler3 struct{}

func (*myHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

// 超时接口
func sayBye3(w http.ResponseWriter, r *http.Request) {
	// 睡眠4秒  上面配置了3秒写超时，所以访问 “/bye“路由会出现没有响应的现象
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye ,this is v3 httpServer"))
}
