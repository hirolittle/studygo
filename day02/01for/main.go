package main

import "fmt"

// 流程控制之跳出 for 循环

func main() {
	for i := 0; i < 10; i++ {
		// 当 i=5 时，跳出 for 循环
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Over")

	// 当 i=5 时，跳过此次 for 循环（不执行 for 循环内部的执行语句），继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("Over")
}
