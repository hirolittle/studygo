package main

import "fmt"

// if 条件判断

func main() {
	age := 19

	if age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("中年了")
	} else if age > 18 {
		fmt.Println("青年了")
	} else {
		fmt.Println("少年了")
	}

	// 作用域
	// age2 变量此时只在 if 判断语句中生效
	if age2 := 19; age2 > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	//undefined: age2
	//fmt.Println(age2) // 在这里找不到 age2
}
