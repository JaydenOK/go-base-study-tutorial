package _func

import (
	"fmt"
	"time"
)

func PrintBr() {
	fmt.Println("###########################################")
}

//
//函数式选项
//这其实也属于可变参数，不过实现思路不同，所以把它分开来了。
//这种方案的核心思路是自定义一个func（type Option func(option *ConnOption)），func的入参是一个指向可选参数结构体的指针，然后为每个可选参数定义一个函数,入参为参数值,该函数返回前面自定义的func,并在返回函数中
//将参数值赋值给可选参数结构体相应的字段：
var (
	defaultTimeOut  = time.Second * 5
	defaultFailFast = true
	defaultSecurity = true
)

// 自定义一个func
type Option func(option *ConnOption)

// 可选参数函数，返回一个Option类型的方法
func WithTimeOut(value time.Duration) Option {
	return func(option *ConnOption) {
		option.timeOut = value
	}
}

func WithFailFast(value bool) Option {
	return func(option *ConnOption) {
		option.failFast = value
	}
}

func WithSecurity(value bool) Option {
	return func(option *ConnOption) {
		option.security = value
	}
}

// 可选参数结构体
type ConnOption struct {
	timeOut  time.Duration
	failFast bool
	security bool
}

type Conn struct {
	url      string
	timeOut  time.Duration
	failFast bool
	security bool
}

func NewConn(url string, options ...Option) *Conn {
	connOption := &ConnOption{
		timeOut:  defaultTimeOut,
		failFast: defaultFailFast,
		security: defaultSecurity,
	}
	for _, opt := range options { // 遍历传入的可选参数函数并执行
		opt(connOption)
	}
	return &Conn{
		url:      url,
		timeOut:  connOption.timeOut,
		failFast: connOption.failFast,
	}
}

func test() {
	//调用
	NewConn("http://www.peachesTao.com", WithTimeOut(time.Second), WithFailFast(true), WithSecurity(false))
	NewConn("http://www.peachesTao.com", WithFailFast(true), WithSecurity(false))
	NewConn("http://www.peachesTao.com", WithSecurity(false))
}
