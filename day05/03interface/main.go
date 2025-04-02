package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (ca cat) move() {
	fmt.Println("cat move")
}

func (ca cat) eat(name string) {
	fmt.Printf("cat eat %s\n", ca.name)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("chicken move")
}

func (c chicken) eat(name string) {
	fmt.Println("chicken eat", name)
}

func main() {
	// 分为两部分，动态类型和动态值
	var a1 animal
	fmt.Printf("%T\n", a1) // 动态类型：nil  动态值：nil
	var c1 = cat{
		name: "淘气",
		feet: 4,
	}
	a1 = c1
	a1.move()
	a1.eat("fish")
	fmt.Printf("%T\n", a1) // 动态类型：main.cat  动态值：cat{name: "淘气", feet: 4}

	var ch = chicken{
		feet: 2,
	}
	a1 = ch
	a1.move()
	a1.eat("虫子")
	fmt.Printf("%T\n", a1) // 动态类型：main.chicken  动态值：chicken{feet: 2}
}
