package main

import "fmt"

// 结构体模拟实现其它语言中的 “继承”

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

type dog struct {
	feet   uint8
	animal //通过嵌套匿名结构体实现继承
}

func (d dog) wang() {
	fmt.Printf("%s 汪汪汪～\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "aa",
		},
		feet: 4,
	}

	d1.wang()
	d1.move()
}
