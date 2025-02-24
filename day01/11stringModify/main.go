package main

import "fmt"

// 字符串修改

// 字符串其实是不能修改的
// 如果需要修改，需要转换成 []byte 或者 []rune

func main() {

	s1 := "hello"
	//s1[0] = 'l' // cannot assign to s1[0]
	fmt.Println(s1)

	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强行修改成 []rune 切片
	s3[0] = '红'
	fmt.Println(string(s3)) // 把 []rune 切片转换成字符串
	fmt.Printf("%p, %s\n", &s2, s2)
	s2 = string(s3)
	fmt.Printf("%p, %s\n", &s2, s2)

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1: %T, c2: %T\n", c1, c2) // c1: string, c2: int32

	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3: %T, c4: %T\n", c3, c4) // c3: string, c4: uint8

}
