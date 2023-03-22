package servers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

//go语言http
//1、net/http 包提供的 http.ListenAndServe() 方法，在指定的地址监听
//该方法用于在指定的 TCP 网络地址 addr 进行监听，然后调用服务端处理程序来处理传入的连 接请求。该方法有两个参数：第一个参数 addr 即监听地址；
//第二个参数表示服务端处理程序， 通常为空，这意味着服务端调用 http.DefaultServeMux 进行处理，
//而服务端编写的业务逻辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中。
//
//2.处理https请求
//func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error
//
//3.路由
//http.HandleFunc()方法接受两个参数
//第一个参数是HTTP请求的 目标路径"/hello"，该参数值可以是字符串，也可以是字符串形式的正则表达式
//第二个参数指定具体的回调方法，比如helloHandler。
//当我们的程序运行起来后，访问http://localhost:8080/hello ， 程序就会去调用helloHandler()方法中的业务逻辑程序。
//开启一个http服务
func server1() {
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

// 开启第二个http服务
// 自己实现一个简单的ServeMux来取代底层默认的ServeMux对象
func server2() {
	// 创建mux来路由请求
	m := http.NewServeMux()
	// 处理所有请求
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	// 启动服务监听8000端口
	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}
	// 持续处理请求直到错误发生
	log.Fatal(s.ListenAndServe())
}

//学完了net/http和fasthttp两个HTTP协议接口的客户端实现，接下来就要开始Server的开发，不学不知道一学吓一跳，居然这两个库还支持Server的开发，太方便了。
//相比于Java的HTTPServer开发基本上都是使用Spring或者Springboot框架，总是要配置各种配置类，各种handle对象。Golang的Server开发显得非常简单，就是因为特别简单，或者说没有形成特别统一的规范或者框架，我发现了很多实现方式，HTTP协议基于还是net/http和fasthttp，但是handle语法就多种多样了。
//先复习一下：Golang语言HTTP客户端实践、Golang fasthttp实践。
//在Golang语言方面，实现某个功能的库可能会比较多，有机会还是要多跟同行交流，指不定就发现了更好用的库。下面我分享我学到的六种Server开发的实现Demo。

type task struct {
	FunTester string
}

//第一种
//基于net/http实现，这是一种比较基础的，对于接口和handle映射关系处理并不优雅，不推荐使用。
func TestHttpSer(t *testing.T) {
	server := http.Server{
		Addr: ":8001",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Index(r.URL.String(), "test") > 0 {
				fmt.Fprintf(w, "这是net/http创建的server第一种方式")
				return
			}
			fmt.Fprintf(w, task.FunTester)
			return
		}),
	}
	server.ListenAndServe()
	log.Println("开始创建HTTP服务")
}

//第二种
//第二种也是基于net/http，这种编写语法可以很好地解决第一种的问题，handle和path有了类似配置的语法，可读性提高了很多。
type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func TestHttpSer2(t *testing.T) {
	http.Handle("/test", &indexHandler{content: "这是net/http第二种创建服务语法"})
	http.Handle("/", &indexHandler{content: task.FunTester})
	http.ListenAndServe(":8001", nil)
}

//第三种
//第三个基于net/http和github.com/labstack/echo，后者主要提供了Echo对象用来处理各类配置包括接口和handle映射，功能很丰富，可读性最佳。
//func TestHttpSer3(t *testing.T) {
//	app := echo.New()
//	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//		AllowOrigins: []string{"*"},
//		AllowMethods: []string{echo.GET, echo.DELETE, echo.POST, echo.OPTIONS, echo.PUT, echo.HEAD},
//		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
//	}))
//	app.Group("/test")
//	{
//		projectGroup := app.Group("/test")
//		projectGroup.GET("/", PropertyAddHandler)
//	}
//	app.Server.Addr = ":8001"
//	gracehttp.Serve(app.Server)
//}

//第四种
//第四种依然基于net/http实现，引入了github.com/gin-gonic/gin的路由，看起来接口和handle映射关系比较明晰了。
//func TestHttpServer4(t *testing.T) {
//	router := gin.New()
//	api := router.Group("/okreplay/api")
//	{
//		api.POST("/submit", gin.HandlerFunc(func(context *gin.Context) {
//			context.JSON(http.StatusOK, task.FunTester)
//
//		}))
//	}
//	s := &http.Server{
//		Addr:           ":8001",
//		Handler:        router,
//		ReadTimeout:    1000 * time.Second,
//		WriteTimeout:   1000 * time.Second,
//		MaxHeaderBytes: 1 << 20,
//	}
//	s.ListenAndServe()
//}

//第五种
//第五种基于fasthttp开发，使用都是fasthttp提供的API，可读性尚可，handle配置倒是更像Java了。
//func TestFastSer(t *testing.T) {
//	address := ":8001"
//	handler := func(ctx *fasthttp.RequestCtx) {
//		path := string(ctx.Path())
//		switch path {
//		case "/test":
//			ctx.SetBody([]byte("这是fasthttp创建服务的第一种语法"))
//		default:
//			ctx.SetBody([]byte(task.FunTester))
//		}
//	}
//	s := &fasthttp.Server{
//		Handler: handler,
//		Name:    "FunTester server",
//	}
//	if err := s.ListenAndServe(address); err != nil {
//		log.Fatal("error in ListenAndServe", err.Error())
//	}
//}

//第六种
//第六种依然基于fasthttp，用到了github.com/buaazp/fasthttprouter，有点奇怪两个居然不在一个GitHub仓库里。使用语法跟第三种方式有点类似，比较有条理，有利于阅读。
//func TestFastSer2(t *testing.T) {
//	address := ":8001"
//	router := fasthttprouter.New()
//	router.GET("/test", func(ctx *fasthttp.RequestCtx) {
//		ctx.Response.SetBody([]byte("这是fasthttp创建server的第二种语法"))
//	})
//	router.GET("/", func(ctx *fasthttp.RequestCtx) {
//		ctx.Response.SetBody([]byte(task.FunTester))
//	})
//	fasthttp.ListenAndServe(address, router.Handler)
//}
