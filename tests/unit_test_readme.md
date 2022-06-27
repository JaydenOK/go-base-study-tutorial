```go
约束：
使用go自身的单元测试框架testing包来写单元测试有如下约束：

单元测试，要导入 testing 包；
承载测试用例的测试文件，固定以 _test.go结尾；
测试用例函数名称一定要以 Test 开头，同时Test后的第一字母一定要大写，如TestSpec，写成这样Testspec不会被框架识别；
测试用例函数的参数有且只有一个，一般是 t *testing.T。
命令：
测试用例文件使用go test指令来执行，没有也不需要 main() 作为函数入口；

所有在以_test结尾的源码内以Test开头的函数会自动被执行；

测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。

go test常见选项：

选项	用途
-v	显示测试的详细命令
-run	-run regexp，只运行 regexp 匹配的函数
-c	生成 test 的二进制文件
使用示例：

运行整个项目的测试文件：
$ go test
PASS
ok      _/home/wangbm/golang/math   0.003s
只运行某个测试文件（ math_test.go， math.go 是一对，缺一不可，前后顺序可对调）
$ go test math_test.go math.go
ok      command-line-arguments  0.002s
加-v查看详细结果
$ go test math_test.go math.go
=== RUN   TestAdd
    TestAdd: main_test.go:22: the result is ok
    TestAdd: main_test.go:22: the result is ok
    TestAdd: main_test.go:22: the result is ok
    TestAdd: main_test.go:22: the result is ok
    TestAdd: main_test.go:22: the result is ok
--- PASS: TestAdd (0.00s)
PASS
ok      command-line-arguments  0.003s
只测试某个函数，-run 支持正则，如下例子中 TestAdd，如果还有一个测试函数为 TestAdd02，那么它也会被运行。
$ go test -v -run="TestAdd"
=== RUN   TestAdd
    TestAdd: math_test.go:22: the result is ok
    TestAdd: math_test.go:22: the result is ok
    TestAdd: math_test.go:22: the result is ok
    TestAdd: math_test.go:22: the result is ok
    TestAdd: math_test.go:22: the result is ok
--- PASS: TestAdd (0.00s)
PASS
ok      _/home/wangbm/golang/math   0.003s
生成 test 的二进制文件：加 -c 参数

   $ go test -v -run="TestAdd" -c
   $
   $ ls -l
   total 3208
   -rw-r--r-- 1 root root      95 May 25 20:56 math.go
   -rwxr-xr-x 1 root root 3272760 May 25 21:00 math.test
   -rw-r--r-- 1 root root     525 May 25 20:56 math_test.go

```