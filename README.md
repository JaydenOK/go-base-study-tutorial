# go-base-learn
JaydenOK




####### gin框架安装  #######
1，创建项目目录，切换到项目目录  

mkdir D:\www\goweb\blog  
cd D:\www\goweb\blog  
  
2，使用官方最新的module管理项目，初始化，生成go.mod文件  
go mod init

3，设置goland IDE环境，开启modules(设置。setting->Go->Go Modules(vgo)->Enable Go Modules[勾上]->设置代理地址)
Proxy代理地址为:     https://goproxy.cn,direct

4，返回项目命令行: 执行安装gin命令
go get -u github.com/gin-gonic/gin

5，测试，创建main.go，代码如下，再执行run，打开浏览器访问： 127.0.0.1:8080/ping ,输出{"message":"pong"}，即gin框架安装成功



```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
