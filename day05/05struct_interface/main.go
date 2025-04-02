package main

import "fmt"

// 同一个结构体实现多个接口
// 接口还可以嵌套

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// cat 实现了 mover 接口
func (c cat) move() {
	fmt.Printf("cat move %s\n", c.name)
}

// cat 实现了 eater 接口
func (c cat) eat(name string) {
	fmt.Printf("cat eat %s\n", c.name)
}

func main() {

}
