package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
}

// go 语言中函数传参数永远传的是拷贝
func f(x person) {
	x.gender = "female" //修改的是副本的 gender
}

func f2(x *person) {
	fmt.Printf("%p\n", x) // 0x14000092080
	//(*x).gender = "female" // 根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "female" // Go 语言语法糖，自动根据指针找到对应的变量
}

func main() {
	var p person
	p.name = "hiro"
	p.gender = "male"

	f(p)
	fmt.Println(p.gender)

	f2(&p)
	fmt.Printf("%p\n", &p) // 0x14000092080
	fmt.Println(p.gender)

	ss := &p
	fmt.Printf("%p\n", ss) // 0x14000092080

	// 1. 结构体指针1
	var p2 = new(person) // 声明变量
	//(*p2).name = "hiro"
	p2.name = "hiro"        // 初始化
	fmt.Printf("%T\n", p2)  // *main.person
	fmt.Printf("%p\n", p2)  // p2 保存的值就是一个内存地址
	fmt.Printf("%T\n", &p2) // **main.person
	fmt.Printf("%p\n", &p2) // p2 的内存地址

	// 2. 结构体指针2
	// 2.1 key-value 声明变量并初始化
	var p3 = &person{
		name:   "hiro",
		gender: "female",
	}
	fmt.Printf("%#v\n", p3)

	var p5 = &person{
		name: "hiro",
	}
	fmt.Printf("%#v\n", p5)

	// 2.2 使用值列表的形式初始化，值的顺序要喝结构体定义时字段的顺序一致
	var p4 = &person{
		"hiro",
		"male",
	}
	fmt.Printf("%#v\n", p4)
}
