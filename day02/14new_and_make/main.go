package main

import "fmt"

// new 和 make

func main() {
	//var a *int // nil pointer
	//*a = 100 // panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println(*a)

	var a = new(int)
	fmt.Println(*a) // 0
	*a = 100
	fmt.Println(*a) // 100

	//var b map[string]int // nil map
	//b["沙河娜扎"] = 100 // panic: assignment to entry in nil map
	//fmt.Println(b)

	var b map[string]int
	b = make(map[string]int, 10)
	fmt.Println(b) // map[]
	b["沙河娜扎"] = 100
	fmt.Println(b) // map[沙河娜扎:100]

}
