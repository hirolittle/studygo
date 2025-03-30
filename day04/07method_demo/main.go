package main

import "fmt"

// 方法和接收者

type Dog struct {
	name string
}

type person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) Dog {
	return Dog{name}
}

func newPerson(name string, age int) *person {
	return &person{name, age}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型首字母小写表示
func (d Dog) wang() {
	fmt.Printf("%s: 汪汪汪～\n", d.name)
}

// 使用值接收者，传拷贝进去
func (p person) guoNian() {
	p.age++
}

// 使用指针接收者，传内存地址进去
func (p *person) zhenGuoNian() {
	p.age++
}

func main() {

	d1 := newDog("hehe")
	d1.wang()

	p1 := newPerson("hiro", 18)
	fmt.Println(p1.age) // 18
	p1.guoNian()
	fmt.Println(p1.age) // 18
	p1.zhenGuoNian()
	fmt.Println(p1.age) // 19

}
