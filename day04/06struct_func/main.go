package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

type student struct {
	name string
}

// 约定成俗，以 new 开头
// 返回的是结构体还是结构体指针呢？
// 当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newStudent(name string) student {
	return student{
		name: name,
	}
}

func main() {
	p1 := newPerson("John", 20)
	p2 := newPerson("Jack", 21)
	p3 := newPerson("Jane", 22)
	fmt.Println(p1, p2, p3)
	fmt.Printf("%p %p %p\n", p1, p2, p3)

	s := newStudent("Jack")
	fmt.Println(s)

}
