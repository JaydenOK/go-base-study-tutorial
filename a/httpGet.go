package a

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
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
