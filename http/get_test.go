package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

//Go语言中的几种单元测试使用：
//文件名
//*_test.go
//函数名
//必须以Test开头，如TestX..或Test_xxx
//函数参数
//必须是指向testing.T的指针：t *testing.T
//
//go test -v
//选项：
//-v 表示提供冗余打印信息
func Test_get(t *testing.T) {
	apiUrl := "http://www.sina.com"
	// URL param map，使用键值对方式生成查询参数query
	data := url.Values{}
	data.Set("name", "枯藤")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}

	u.RawQuery = data.Encode() //键值对连接, url_encode

	fmt.Println(u.RawQuery) //查询参数query

	fmt.Println(u.String()) //完整url
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

// server端接收
//func getHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	data := r.URL.Query()
//	fmt.Println(data.Get("name"))
//	fmt.Println(data.Get("age"))
//	answer := `{"status": "ok"}`
//	w.Write([]byte(answer))
//}
