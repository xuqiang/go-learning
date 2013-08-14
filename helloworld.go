// 所有的go文件需要以package开头
package main

// 这里是go的注释
/*
 * 多行的注释
 */

// 导入fmt包
import "fmt"

// go运行时，同样是从main函数开始执行的，这里我们定义main
func main() {
	fmt.Printf("hello world")
}

