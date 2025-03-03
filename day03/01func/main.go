package main

import (
	"fmt"
)

// 函数：一段代码的封装

func f1() {
	fmt.Println("hello world")
}

func f2(name string) {
	fmt.Println("hello", name)
}

// 带参数和返回值的函数
func f3(x int, y int) int {
	return x + y
}

// 参数类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(title string, y ...int) int {
	fmt.Println(title, y) // y 是 int 类型的切片
	return 0
}

// 命名返回值
func f6(x, y int) (sum int) {
	sum = x + y // 如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return      // 如果使用命名的返回值，return 后面可以省略返回值变量
}

// Go 语言支持多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("hiro")
	fmt.Println(f3(1, 2))
	fmt.Println(f4(1, 2))
	fmt.Println(f5("hello", 1, 2, 3))
	fmt.Println(f6(1, 2))
	fmt.Println(f7(1, 2))

	// 在一个命名的函数中，不能够再次声明命名函数
	//func f8() {
	//	fmt.Println("f8")
	//}
}
