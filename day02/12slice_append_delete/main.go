package main

import "fmt"

// 关于 append 删除切片中的某个元素

func main() {

	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	s := a[:]

	// 删除索引为 1 的那个 3
	s = append(s[:1], s[2:]...)
	fmt.Println(s)
	fmt.Println(a)
}
