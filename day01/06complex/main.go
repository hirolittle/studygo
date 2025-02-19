package main

import "fmt"

// 复数
// 了解即可

func main() {
	var c1 complex64
	c1 = 1 + 2i
	fmt.Printf("%v, %T\n", c1, c1)

	var c2 complex128
	c2 = 3 + 4i
	fmt.Printf("%v, %T\n", c2, c2)
}
