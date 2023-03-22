package utils

import (
	"encoding/json"
	"net/http"
)

//获取参数方式
//id := query["id"][0]
//id := query.Get("id")
func ParamsGet(req *http.Request) interface{} {
	query := req.URL.Query()
	return query
}

func ParamsPost(req *http.Request) interface{} {
	req.ParseForm()
	return req.PostForm
}

func ParamsJson(req *http.Request) interface{} {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(req.Body)
	// params用于存放参数key=value数据，或解析到结构体
	var params map[string]string
	// 解析参数 存入map
	decoder.Decode(&params)
	return params
}

func RequestHeader(req *http.Request) interface{} {
	return nil
}

func HttpMethod(req *http.Request) interface{} {
	return req.Method
}
