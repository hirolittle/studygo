package main

import "fmt"

// 自定义类型和类型别名
// type 后面跟的是类型

// 自定义类型
type myInt int

// 类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Printf("%d, %T\n", n, n) // 100, main.myInt

	var m yourInt
	m = 100
	fmt.Printf("%d, %T\n", m, m) // 100, int

	var x int
	x = 100
	fmt.Printf("%d, %T\n", x, x) // 100, int

	var c rune
	c = '中'
	fmt.Printf("%v, %T\n", c, c) // 20013, int32

}
