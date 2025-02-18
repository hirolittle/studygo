package main

// 导入语句
import "fmt"

// 函数外只能放置标识符（变量、常量、函数、类型）的声明

// 声明变量
var name string
var age int
var isOK bool

// 批量声明
//var (
//	name string // ""
//	age  int    // 0
//	isOK bool   // false
//)

// 程序入口
func main() {

	name = "hiro"
	age = 18
	isOK = true

	fmt.Printf("name: %s, age: %d, isOK: %t\n", name, age, isOK)

	fmt.Println("Hello, world!")

	// 非全局变量声明后，即使赋值，也是未使用：declared and not used: aa
	//var aa int
	//aa = 1

	// 声明并同时赋值
	var s1 string = "hello"
	fmt.Println(s1)

	// 类型推导，根据值判断该变量是什么类型
	var s2 = "world"
	fmt.Println(s2)

	// 简短变量声明，只能用于函数内部
	s3 := "golang"
	fmt.Println(s3)

	// 匿名变量
	// _

	// 同一个作用域中，不能重复声明同一个变量
	//s1 := "hello"
}
