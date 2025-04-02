package main

import "fmt"

// 空接口

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("type: %T, value: %v\n", a, a)
}

// interface: 关键字
// interface{}: 空接口
func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "hiro"
	m1["age"] = 18
	m1["married"] = true
	m1["hobby"] = [...]string{"hanako", "football"}

	fmt.Printf("%#v\n", m1)

	fmt.Printf("%T\n", m1["name"])

	show(m1)
	show(false)
	var e interface{}
	show(e)

}
