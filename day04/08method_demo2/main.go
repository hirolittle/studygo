package main

import "fmt"

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己的包里

// MyInt 将int定义为自定义MyInt类型
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (i MyInt) SayHello() {
	fmt.Println("我是一个 int")
}

func main() {

	var m MyInt
	m = 100
	fmt.Printf("%v %T\n", m, m) // 100 main.MyInt
	m.SayHello()

}
