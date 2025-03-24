package main

import "fmt"

// 结构体占用一块连续的内存

type x struct {
	a int8 // 8bit -> 1byte
	b int8
	c int8
	d int8
}

func main() {
	m := x{
		a: 10,
		b: 20,
		c: 30,
		d: 40,
	}

	fmt.Printf("%p\n", &(m.a)) // 0x1400010200c
	fmt.Printf("%p\n", &(m.b)) // 0x1400010200d
	fmt.Printf("%p\n", &(m.c)) // 0x1400010200e
	fmt.Printf("%p\n", &(m.d)) // 0x1400010200f
}
