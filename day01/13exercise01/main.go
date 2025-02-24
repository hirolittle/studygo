package main

import "fmt"

// 编写代码，分别定义一个整型、浮点型、布尔型、字符串变量
// 使用 fmt.Printf() 搭配 %T 分别打印出上述变量的值和类型

func main() {

	i := 10
	f := 10.1
	var b bool
	s := "hello"

	fmt.Printf("%d, %T\n", i, i)
	fmt.Printf("%f, %T\n", f, f)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", s, s)

}
