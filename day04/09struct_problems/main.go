package main

import "fmt"

// 结构体遇到的问题

type MyInt int

func (i MyInt) SayHello() {
	fmt.Println("我是一个 int")
}

type person struct {
	name string
	age  int
}

func main() {
	// 问题1： MyInt(100) 是个啥

	// 声明一个 int32 类型的变量 x， 它的值是 10
	// 方法1
	//var x int32
	//x = 10
	// 方法2
	//var x int32 = 10
	// 方法3
	//var x = int32(10)
	// 方法4
	x := int32(10)
	fmt.Printf("%d, %T\n", x, x) // 10, int32

	// 声明一个 MyInt 类型的变量 m， 它的值是 100
	m := MyInt(100)
	fmt.Printf("%d, %T\n", m, m) // 100, main.MyInt
	m.SayHello()

	// 问题2：结构体初始化

	// 方法1
	var p person
	p.name = "Jack"
	p.age = 20
	fmt.Println(p)

	// 方法2 键值对初始化
	p2 := person{
		name: "Jack",
		age:  20,
	}
	fmt.Println(p2)

	// 方法3 值列表初始化
	p3 := person{"Jack", 20}
	fmt.Println(p3)

}

// 问题3：为什么要有构造函数
// 可以想到用函数做封装初始化重复代码么
func newPerson(name string, age int) *person {
	// 别人调用我，我能给他一个 person 类型的变量
	return &person{
		name: name,
		age:  age,
	}
}
