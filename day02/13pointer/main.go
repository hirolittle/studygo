package main

import "fmt"

// pointer

func main() {
	// 1. & 取地址
	// 2. * 根据地址取值

	num := 10
	fmt.Println(&num)

	p := &num
	fmt.Println(p)
	fmt.Printf("%T\n", p) // *int: int 类型的指针

	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) // int

}
