package http

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_file(t *testing.T) {
	fileName := "../logs/a.txt"
	//bytes, _ := readFile(fileName)
	//fmt.Println(string(bytes))

	str := `交互模式
	请求：Http Post请求跳转交互模式 返回结果+通知
	测试地址：https://demoapi.yl-scm.com/yunlu-order-web/inventory/inventoryAction!inventoryQuery.action`

	size, _ := writeFile(fileName, str)
	fmt.Println("写入字节数:", size)
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("打开文件错误：", err.Error())
		return nil, err
	}
	defer file.Close()
	//每次读取1024字节到buffer
	b := make([]byte, 1024)
	var content []byte
	for {
		//在文件末尾，Read返回0,io.EOF
		size, err := file.Read(b)
		if err == io.EOF {
			break
		}
		buffer := b[:size]
		//使用 ... 运算符将一个切片添加到另一个切片
		content = append(content, buffer...)
	}
	return content, nil
}

// 写入文件（追加文件内容）
func writeFile(fileName string, content string) (int, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("打开文件错误：", err.Error())
		return 0, err
	}
	defer file.Close()
	b := []byte(content)
	//它返回写入的字节数和错误（如果有）
	size, err := file.Write(b)
	//size, err := file.WriteString(content)		//直接写入字符串
	if err != nil {
		fmt.Println("写入文件错误：", err.Error())
		return 0, err
	}
	// 确保写入到磁盘
	err = file.Sync()
	if err != nil {
		return 0, err
	}
	return size, nil
}

// 追加文件内容
func appendFile() {
	//file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

// 三、判断文件是否存在
func test6() {
	file := "testFile"
	_, err := os.Stat(file)
	fmt.Println(!os.IsNotExist(err))
}

// 四、文件拷贝
func test7() {
	srcFile, _ := os.Open("testFile")
	defer srcFile.Close()
	desFile, _ := os.OpenFile("copyFile", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer desFile.Close()

	io.Copy(desFile, srcFile)
}

// 五、改变程序当前工作目录
func test8() {
	fmt.Println(os.Getwd())
	err := os.Chdir("..")
	if err != nil {
		return
	}
	fmt.Println(os.Getwd())
}
