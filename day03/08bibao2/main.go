package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了它外部作用域的一个变量
// 闭包 = 函数 + 外部变量的引用

// 底层的原理：
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序，先在自己的内部找，找不到往外层找

func adder1() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret1 := adder1()
	fmt.Println(ret1(200))
	fmt.Println(ret1(200))

	ret2 := adder2(100)
	fmt.Println(ret2(200))
	fmt.Println(ret2(200))
}
