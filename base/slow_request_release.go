package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Go1.7加入了一个新的标准库context，它定义了Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作， 这些操作可能涉及多个 API 调用。
// 对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用
// WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

// 客户端
type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}
	url := "http://127.0.0.1:8000/"
	url = "https://www.baidu.com/"
	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) // 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()
	//然后使用一个select让主程序陷入等待
	select {
	case <-ctx.Done():
		//transport.CancelRequest(req)
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}
}

//============ 使用Context的注意事项
//推荐以参数的方式显示传递Context
//以Context作为参数的函数方法，应该把Context作为第一个参数。
//给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TOD0
//Context的Value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
//Context是线程安全的，可以放心的在多个goroutine中传递

func main() {
	// 定义一个200毫秒的超时,WithTimeout() 得到一个上下文（ctx）和一个取消函数（cancel）
	ctx, cancelA123 := context.WithTimeout(context.Background(), time.Millisecond*200)
	//取消此上下文将释放与其相关的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel，通常用于数据库或者网络连接的超时控制。
	defer cancelA123() // 调用cancel释放子goroutine资源
	doCall(ctx)
}
