package main

import "fmt"

// 切片 slice

func main() {

	// 切片的定义
	var s1 []int    // 定义一个存放 int 类型元素的切片
	var s2 []string // 定义一个存放 string 类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d, cap(s2): %d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 基于一个数组切割，左包含右不包含，左闭右开
	fmt.Println(s3, len(s3), cap(s3))
	s4 := a1[1:6]
	fmt.Println(s4, len(s4), cap(s4))
	s5 := a1[:4] // => [0:4] [1 3 5 7]
	s6 := a1[3:] // => [3:len(a1)] [7 9 11 13]
	s7 := a1[:]  // => [0:len(a1)]
	fmt.Println(s5, s6, s7)

	// 切片的容量是指底层数组的容量，底层数组从切片的第一个元素到最后的元素数量
	fmt.Println(s5, len(s5), cap(s5))
	fmt.Println(s6, len(s6), cap(s6))

	// 切片再切片
	s8 := s5[3:7]
	fmt.Println(s8, len(s8), cap(s8))

	// 切片是引用类型，都指向了底层的一个数组
	a1[6] = 1300 // 修改了底层数组的值
	fmt.Println(s6)
	fmt.Println(s8)

}
