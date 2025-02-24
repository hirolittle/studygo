package main

import "fmt"

// for 循环

func main() {

	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 初始语句可以被忽略，但是初始语句后的分号必须要写
	j := 10
	for ; j < 20; j++ {
		fmt.Println(j)
	}

	// 初始语句和结束语句都可以省略，类似于其他编程语言中的`while`
	k := 5
	for k < 10 {
		fmt.Println(k)
		k++
	}

	// 无限循环，死循环
	//for {
	//	fmt.Println("infinite loop")
	//}

	// for range(键值循环)
	s := "hello世界"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
