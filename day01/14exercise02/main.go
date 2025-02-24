package main

import "fmt"

// 编写代码统计出字符串 "hello世界" 中汉字的数量

func main() {
	s := "hello世界"
	count := 0
	for _, v := range s {
		if v >= 0x4e00 && v <= 0x9fa5 {
			count++
		}
	}

	fmt.Println(count)
}
