package main

import "fmt"

// 函数

// 函数存在的定义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接用函数名调用就可以了
// 使用函数能够使代码结构更清晰、更简洁

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数
func f2() {
	fmt.Println("hello")
}

// 没有参数，但有返回值的
func f3() int {
	ret := 100
	return ret
}

// 参数可以命名，也可以不命名
// 命名的返回值，就相当于在函数中声明了一个变量
func f4(x, y int) (ret int) {
	ret = x + y
	return // 使用命名返回值，return 可以省略
}

// 多个返回值
func f5() (int, int) {
	return 1, 2
}

// 参数类型简写
// 当参数中连续两个参数的类型一致时，我们可以将前面那个参数的类型省略
func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 的类型是切片 []int
}

// Go 语言中函数没有默认参数这个概念，设计哲学：所见即所得

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	f7("hello", 1, 2, 3)
}
