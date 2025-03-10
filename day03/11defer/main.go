package main

import "fmt"

// defer

// 提示：
// defer 注册要延迟执行的函数时该函数所有的参数都需要确定其值

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	//defer calc("AA", x, y)
	x = 10
	defer calc("BB", x, calc("B", x, y))
	//defer calc("BB", x, y)
	y = 20
}

// 1. defer calc("AA", x, calc("A", x, y))
// 2. defer calc("AA", 1, 3) // "A" 1 2 3
// 3. x =  10
// 4. defer calc("BB", 10, 12) // "B" 10 2 12
// 5. y = 20
// 6. calc("BB", 10, 12) // "BB" 10 12 22
// 7. calc("AA", 1, 3) // "AA" 1 3 4
