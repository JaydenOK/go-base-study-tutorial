# go-base-learn-tutorial
go(golang)入门教程整理，go常用工具类封装、整理  
golang安装，模块化设置，gin框架安装  

### Linux golang安装
```
1.下载golang
golang所有版本网址
https://studygolang.com/dl

//下载并解压到/usr/local文件下
wget https://studygolang.com/dl/golang/go1.19.1.linux-amd64.tar.gz
//解压并复制到/user/local文件夹下
tar -C /usr/local -zxf go1.19.1.linux-amd64.tar.gz

2.编辑环境变量文件
vim /etc/profile

在文件后追加以下内容
export GOPROXY=https://goproxy.cn
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/root/go
export PATH=$PATH:$GOPATH/BIN

退出并保存，刷新环境变量
source /etc/profile

3.检验安装是否成功
go version

4. 环境变量信息
 go env

5. 卸载
rm -rf /usr/local/go

-- (windows安装)
Windows下可以使用.msi 后缀(如go1.19.1.windows-amd64.msi)的安装包来安装。比较简单
```

### 项目模块化module设置
```text
1，创建项目目录，切换到项目目录  

mkdir D:\www\goweb\mall  
cd D:\www\goweb\mall  
  
2，使用官方最新的module管理项目，初始化，生成go.mod文件  
go mod init mall

3，然后修改golang IDE设置
run -> Edit Configurations -> Working directory修改为 
  
4，设置goland IDE环境，开启modules(设置。setting->Go->Go Modules(vgo)->Enable Go Modules[勾上]->设置代理地址) (一定要勾选)
然后设置Proxy代理地址为:     https://goproxy.cn,direct

5，新增 main.go 文件， 设置 package main

6，新增api目录及文件，包名package命名为当前目录名（其它目录也按此规则设置包名）
mkdir D:\www\goweb\api
touch D:\www\goweb\api\user.go

package api

import (
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	c.JSON(200, "OK")
}
```


### gin框架安装
```text
1, 创建项目目录，切换到项目目录  

mkdir D:\www\goweb\blog  
cd D:\www\goweb\blog  
  
2，使用官方最新的module管理项目，初始化，生成go.mod文件  
go mod init

3，设置goland IDE环境，开启modules(设置。setting->Go->Go Modules(vgo)->Enable Go Modules[勾上]->设置代理地址)
Proxy代理地址为:     https://goproxy.cn,direct

4，返回项目命令行: 执行安装gin命令
go get -u github.com/gin-gonic/gin

5，测试，创建main.go，代码如下，再执行run，打开浏览器访问： 127.0.0.1:8080/ping ,输出{"message":"pong"}，即gin框架安装成功


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