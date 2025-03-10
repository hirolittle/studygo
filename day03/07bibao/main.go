package main

import "fmt"

// 闭包

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 要求：
// f1 是其他人定义的函数
// f2 是自己写的函数
// 想把自己写的函数 f2 传给 f1，执行一下，类似这样 f1(f2)
// f2 直接传给 f1，函数类型不同，无法直接传
// 定义一个函数对 f2 进行包装
func f3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func main() {

	// 实现 f1(f2)
	f := f3(f2, 1, 2) // 把原来需要传递两个 int 类型的参数包装成一个不需要传参的函数
	f1(f)
}
