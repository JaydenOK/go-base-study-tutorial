package servers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

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

func httpGet() {
	apiUrl := "http://rest.java.yibainetwork.com/distribution/dcmShop/getDcmShopList"
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

func httpPost() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=枯藤&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"枯藤","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}

	fmt.Println(string(b))
}

//server 端接收
//func postHandler(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()
//	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
//	r.ParseForm()
//	fmt.Println(r.PostForm) // 打印form数据
//	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
//	// 2. 请求类型是application/json时从r.Body读取数据
//	b, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println("read request.Body failed, err:%v\n", err)
//		return
//	}
//	fmt.Println(string(b))
//	answer := `{"status": "ok"}`
//	w.Write([]byte(answer))
//}
