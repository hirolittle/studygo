package main

import "fmt"

// fmt 占位符

func main() {

	n := 100
	s := "hello"

	// 查看类型
	fmt.Printf("%T, %T\n", n, s)

	// 查看值
	fmt.Printf("%v, %v\n", n, s)
	fmt.Printf("%+v, %+v\n", n, s)
	fmt.Printf("%#v, %#v\n", n, s)
	fmt.Printf("%d, %s\n", n, s)

	// 进制表示
	fmt.Printf("%b, %d, %o, %x\n", n, n, n, n)

}
