package main

import "fmt"

// defer

// 多用于函数结束之前释放资源（文件句柄、数据库连接、socket 连接）
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer 把后面的语句延迟到函数返回的时候再执行
	defer fmt.Println("呵呵呵") // 一个函数可以有多个 defer 语句
	defer fmt.Println("哈哈哈") // 多个 defer 语句，按照先进后出(后进先出)的顺序延迟执行
	fmt.Println("end")
}

func main() {
	deferDemo()
}
