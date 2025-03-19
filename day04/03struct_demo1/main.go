package main

import "fmt"

// 结构体

type person struct {
	name    string
	age     int
	gender  string
	hobbies []string
}

func main() {

	// 声明一个 person 类型的变量 p
	var p person

	// 通过字段赋值
	p.name = "Jack"
	p.age = 23
	p.gender = "male"
	p.hobbies = []string{"python", "golang"}

	fmt.Printf("%T\n", p) // main.person

	fmt.Println(p)         // {Jack 23 male [python golang]}
	fmt.Printf("%v\n", p)  // {Jack 23 male [python golang]}
	fmt.Printf("%+v\n", p) // {name:Jack age:23 gender:male hobbies:[python golang]}
	fmt.Printf("%#v\n", p) // main.person{name:"Jack", age:23, gender:"male", hobbies:[]string{"python", "golang"}}

	// 访问变量 p 的字段
	fmt.Println(p.name)

	// 匿名结构体
	// 一些临时数据结构等场景下
	var user struct {
		name string
		age  int
	}
	user.name = "Jack"
	user.age = 23
	fmt.Printf("%T\n", user)  // struct { name string; age int }
	fmt.Printf("%#v\n", user) // struct { name string; age int }{name:"Jack", age:23}
}
