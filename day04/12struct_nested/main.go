package main

import "fmt"

// 嵌套结构体

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name    string
	age     int
	address address
}

type company struct {
	name    string
	address // 匿名嵌套结构体
}

type employee struct {
	name string
	address
	workPlace
}

func main() {
	p1 := person{
		name: "hiro",
		age:  22,
		address: address{
			province: "zhejiang",
			city:     "hangzhou",
		},
	}

	fmt.Printf("%#v\n", p1)
	// main.person{name:"hiro", age:22, address:main.address{province:"zhejiang", city:"hangzhou"}}

	fmt.Println(p1.age, p1.address.city)

	c1 := company{
		name: "huawei",
		address: address{
			province: "zhejiang",
			city:     "hangzhou",
		},
	}
	fmt.Printf("%#v\n", c1)
	fmt.Println(c1.address.city) // 匿名字段默认使用类型名作为字段名
	// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找
	fmt.Println(c1.city) // 匿名字段可以省略

	e1 := employee{
		name: "hiro",
		address: address{
			province: "zhejiang",
			city:     "hangzhou",
		},
		workPlace: workPlace{
			province: "zhejiang",
			city:     "hangzhou",
		},
	}
	fmt.Printf("%#v\n", e1)
	//fmt.Println(e1.province) // Ambiguous reference 'province'
	fmt.Println(e1.address.city)
	fmt.Println(e1.workPlace.city)
}
