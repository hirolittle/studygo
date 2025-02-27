package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	a3 := make([]int, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将 a1 中的索引为 1 的 3 这个元素删掉
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, cap(a1))

	x1 := [...]int{1, 3, 5} // 数组
	fmt.Println(x1)
	s1 := x1[:] // 切片
	fmt.Println(s1, len(s1), cap(s1))

	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是占用一块连续的内存
	fmt.Printf("%p\n", &s1[0])     // 0x140000a8030
	s1 = append(s1[:1], s1[2:]...) // 修改了底层数组!!!
	fmt.Printf("%p\n", &s1[0])
	fmt.Println(s1, len(s1), cap(s1)) // 0x140000a8030
	fmt.Println(x1)                   // [1 5 5]

	s1[0] = 100
	fmt.Println(x1) // [100 5 5]
}
