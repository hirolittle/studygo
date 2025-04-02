package main

import "fmt"

// 使用值接收者和指针接收者的区别

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
func (c cat) move() {
	fmt.Println("cat move")
}

func (c cat) eat(name string) {
	fmt.Printf("cat eat %s\n", c.name)
}

type dog struct {
	name string
	feet int8
}

func (d *dog) move() {
	fmt.Println("dog move")
}

func (d *dog) eat(name string) {
	fmt.Printf("dog eat %s\n", d.name)
}

func main() {
	var a1 animal
	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"jim", 4} // *cat

	a1 = c1
	a1.move()
	fmt.Printf("type: %T, a1=%#v\n", a1, a1) // type: main.cat, a1=main.cat{name:"tom", feet:4}
	// fmt.Printf 会调用 fmt 包的反射机制（reflect），它能够获取接口的 动态类型 和 动态值
	// 虽然 fmt.Println(a1) 可以显示 a1 的值，但 a1 仍然是 animal 类型，而 接口类型本身不包含具体类型的字段。
	// 换句话说，animal 接口只定义了 move() 和 eat() 方法，它并不知道具体的 cat 结构体里面有 name 字段
	// 要访问 cat 结构体的字段，就必须进行 类型断言
	if c, ok := a1.(cat); ok {
		fmt.Printf("type: %T, c=%#v\n", c, c)
		fmt.Println(c.name)
	}

	a1 = c2 // Go 自动对指针求值，相当于 a1 = *c2，所以 move() 方法仍然匹配接口
	a1.move()
	fmt.Println(a1)

	var a2 animal
	//d1 := dog{"tom", 8}
	d2 := &dog{"jim", 4}

	// Go 允许自动对指针求值，但 不会自动对值取地址，需要手动 &d1 取地址
	//a2 = d1 // Type does not implement animal as the move method has a pointer receiver
	//a2.move()

	a2 = d2
	a2.move()

}
