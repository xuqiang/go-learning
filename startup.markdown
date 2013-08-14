# sartup
## go一些特性
1. Concurrent
 Go makes it easy to “ﬁre off” functions to be run as very lightweight threads. These
 threads are called goroutines a in Go;
2. channels
 go内部的goroutines之间使用channel交互
...

## 安装
官网上下载系统对应版本的go二进制发行版本。添加环境变量，如下
	export GOROOT=$HOME/go/
	export PATH=$PATH:$GOROOT/bin/ 
将go添加到path环境变量中

## hello world using go
代码及注释参考这里：https://raw.github.com/xuqiang/go-learning/master/helloworld.go 。首先build代码
go build helloworld.go
此时将生成helloworld二进制，执行即可。或者直接使用go run helloworld.go

## Variables, types and keywords变量 数据类型 关键字
1. go里面一行可以写多个代码，需要使用;分割
2. go中定义变量的语法和通常c不同，c语言中我们使用int a; 但是在go中需要使用 var a int. 结合赋值语句为
var a int = 5
如果仅仅是声明(declare)了变量，go默认会附初始值.例如
	var a int
那么a默认将是0
3. go中_的下划线赋予的值将会被丢弃，例如
	_, b = 10, 11
那么b的值将为11
4. 如果是定义declare但是未使用的变量，在go中会出现编译error "a declared and not used"

