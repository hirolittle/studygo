package main

import "fmt"

// 引出接口的实例

// 定义一个能叫的类型
type speaker interface {
	speak() // 只要实现了 speak 方法的变量，都是 speaker 类型
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪～")
}

func (p person) speak() {
	fmt.Println("啊啊啊～")
}

// 接收一个参数，传进来什么，我就打什么
func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)
}
