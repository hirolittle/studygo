package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// 本来 \ 是有特殊含义的，我应该告诉程序，这是一个单纯的反斜线
	path := "D:\\go\\src\\github.com\\go-programming-tour-book"
	fmt.Println(path)

	// 多行字符串
	s1 := `
	123
	456
	`
	fmt.Println(s1)

	s2 := `D:\go\src\github.com\go-programming-tour-book`
	fmt.Println(s2)

	// 字符串相关操作
	// 字符串长度
	fmt.Println(len(s2))

	// 字符串拼接
	name := "hiro"
	hobby := "coding"
	s3 := name + " " + hobby
	s4 := fmt.Sprintf("%s %s", name, hobby)
	fmt.Printf("%s %s\n", name, hobby)
	fmt.Println(s3)
	fmt.Println(s4)

	// 字符串分割
	split := strings.Split(path, "\\")
	fmt.Println(split)

	// 字符串包含
	fmt.Println(strings.Contains(path, "com"))

	// 字符串前缀
	fmt.Println(strings.HasPrefix(path, "D:"))

	// 字符串后缀
	fmt.Println(strings.HasSuffix(path, "src"))

	// 字符串子串索引
	fmt.Println(strings.Index(path, "go"))
	fmt.Println(strings.LastIndex(path, "go"))

	// 字符串拼接
	fmt.Println(strings.Join(split, "_"))
}
