package main

import "fmt"

// 类型断言

func assign(a interface{}) {
	//fmt.Printf("%T\n", a)
	s, ok := a.(string)
	if !ok {
		fmt.Println("not string")
		return
	}
	fmt.Printf("string, value is %s\n", s)
}

func assign2(a interface{}) {
	//fmt.Printf("%T\n", a)
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	assign(10)
	assign("b")

	assign2(10)
	assign2("b")
	assign2(true)
}
