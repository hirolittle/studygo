package main

import "fmt"

// 接口的定义与实现

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 speak 方法的变量，都是 speaker 类型
}

type cat struct{}

type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪～")
}

// 接收一个参数，传进来什么，我就打什么
func da(x speaker) {
	x.speak()
}

func main() {

	var c1 cat
	var d1 dog

	var s1 speaker // 定义一个接口类型 speaker 变量 s1
	s1 = c1
	fmt.Printf("%T\n", s1) // main.cat
	da(s1)

	s1 = d1
	fmt.Printf("%T\n", s1) // main.dog
	da(s1)
}
