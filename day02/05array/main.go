package main

import "fmt"

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）

func main() {

	// 数组声明
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1: %T, a2: %T\n", a1, a2) // a1: [3]bool, a2: [4]bool

	// 数组初始化
	// 如果不初始化，默认元素都是零值（bool：false, int/float: 0, string: ""）
	fmt.Println(a1, a2)

	// 初始化方式 1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	// 初始化方式 2
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)

	// 初始化方式 3
	// 根据索引来初始化
	a3 := [5]int{1, 2}
	fmt.Println(a3)
	a31 := [5]int{0: 1, 4: 2}
	fmt.Println(a31)

	// 数组遍历
	citys := [...]string{"北京", "上海", "深圳"}

	for i := 0; i < len(citys); i++ {
		fmt.Println(i, citys[i])
	}

	for i, city := range citys {
		fmt.Println(i, city)
	}

	// 多维属组
	// [[1 2] [3 4] [5 6]]
	var a4 [3][2]int
	a4 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a4)

	for _, v1 := range a4 {
		for _, v2 := range v1 {
			fmt.Printf("%d \n", v2)
		}
	}

	// 数组是值类型，类比文件的赋值拷贝
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[2] = 100
	fmt.Println(b1, b2)
}
