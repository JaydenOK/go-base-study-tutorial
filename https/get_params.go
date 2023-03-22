package https

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//golang获取postman传递数据的方法
//http.request的三个属性Form、PostForm、MultipartForm:
//Form：存储了post、put和get参数，在使用之前需要调用ParseForm方法。
//PostForm：存储了post、put参数，在使用之前需要调用ParseForm方法。
//MultipartForm：存储了包含了文件上传的表单的post参数，在使用前需要调用ParseMultipartForm方法。

//获取Get参数
//用postman测试，提交，服务端输出 ：[111]，提交： ;uid=222。服务端输出：[111 222]
//小结：r.Form是url.Values字典类型，r.Form[“id”]取到的是一个数组类型。因为http.request在解析参数的时候会将同名的参数都放进同一个数组里。
//Golang中如何处理POST上来的数组数据
//如上代码执行结果如下
//POST
//map[]

//通过输出body可以看出，业务服务器发送过来的数据是标准的post的数据，可是为什么无论用什么方法取出来的都是空呢？如何是php可以用$_POST,那么golang要用什么方法获取post过来的数据呢
//关于post参数如何提取
//通过php，jsp等语言 把post方式传过来的值赋给隐藏标签的value属性
//如：input id='dd' type='hidden' value='? echo $_POST[参数名];?'/input
//然后在该标签的后边加入你想对传过来的参数操作的js代码
//js中获取值的代码：document.getElementById('dd').value

func getParams(w http.ResponseWriter, r *http.Request) {
	//获取所有请求参数
	query := r.URL.Query()
	jsonStr, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Fprintf(w, "get all param: "+string(jsonStr)+"\n")
	//获取指定请求参数
	names, ok := query["name"]
	if !ok || len(names[0]) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}
}

// 处理application/json类型的POST请求
func handlePostJson(writer http.ResponseWriter, request *http.Request) {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(request.Body)
	// 用于存放参数key=value数据
	var params map[string]string
	// 解析参数 存入map
	decoder.Decode(&params)
	fmt.Printf("POST json: username=%s, password=%s\n", params["username"], params["password"])
	fmt.Fprintf(writer, `{"code":0}`)
}