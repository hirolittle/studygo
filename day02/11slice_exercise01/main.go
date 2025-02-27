package main

import (
	"fmt"
	"sort"
)

// 切片的练习

func main() {

	var b []string // []
	for i := 0; i < 10; i++ {
		b = append(b, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%#v\n", b) // []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var a = make([]string, 5, 10) // ['' '' '' '' '']
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%#v\n", a) // []string{"", "", "", "", "", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	a1 := [...]int{3, 7, 1, 9, 8}
	sort.Ints(a1[:]) // 对切片进行排序
	fmt.Println(a1)
}
