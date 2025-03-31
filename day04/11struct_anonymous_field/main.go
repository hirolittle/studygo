package main

import "fmt"

// 匿名字段
// 字段比较少比较简单的场景
// 不常用

type person struct {
	string
	int
}

func main() {
	p1 := person{"hiro", 18}
	fmt.Printf("p1 type is %T\n", p1) // p1 type is main.person
	fmt.Println(p1.string, p1.int)    // hiro 18
}
