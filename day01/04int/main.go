package main

import "fmt"

// 整型

func main() {

	// 十进制数
	var i1 = 10
	fmt.Printf("i1: %d\n", i1)

	// 八进制数
	i2 := 077
	fmt.Printf("i2: %o, %O, %d\n", i2, i2, i2)

	// 十六进制数
	i3 := 0xFF
	fmt.Printf("i3: %x, %X, %d\n", i3, i3, i3)

	// 二进制数
	i4 := 0b1010
	fmt.Printf("i4: %b, %d\n", i4, i4)

	// 用 _ 来分隔数字
	i5 := 1_000_000
	fmt.Printf("i5: %d\n", i5)

	// 十六进制，后面 pn，表示 2 的 n 次方
	i6 := 0x1p10
	fmt.Printf("i6: %x, %X, %f\n", i6, i6, i6)

	// 查看类型
	i7 := 100
	fmt.Printf("i7: %T\n", i7)

	// 声明一个 int8 类型变量
	i8 := int8(100)
	fmt.Printf("i8: %d, %T\n", i8, i8)

}
