package main

import "fmt"

// byte 和 rune 类型

// byte 型，代表一个 ASCII码 字符，实际是 uint8 类型的别名
// rune 类型，Go 语言中为了非 ASCII 码类型的字符，定义了新的类型，实际是 int32 类型的别名

func main() {
	s := "hello世界"
	// len() 求得是 byte 字节数组的长度
	n := len(s)
	fmt.Println(n)

	for i := 0; i < n; i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	fmt.Println()

	// 从字符串中拿出具体的字符
	for _, v := range s {
		fmt.Printf("%v(%c) ", v, v)
	}
}
