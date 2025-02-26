package main

import "fmt"

// 运算符

func main() {

	// 算术运算符
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ // 单独的语句，不能在等号的右边去赋值 a = a + 1
	b-- // 单独的语句，不能在等号的右边去赋值 b = b - 1
	fmt.Println(a, b)

	// 关系运算符
	fmt.Println(a == b) // Go 语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)

	// 逻辑运算符
	// &&, ||, !
	age := 22

	// &&，and, 如果年龄大于18岁 并且 年龄小于60岁
	if age > 18 && age < 66 {
		fmt.Println("上班族")
	} else {
		fmt.Println("不用上班")
	}

	// ||, or, 如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 66 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班族")
	}

	// !, not, 取反，原来为真就为假，原来为假就为真
	isMale := true
	fmt.Println(!isMale)

	// 位运算符：针对的是二进制数
	// 5 的二进制表示：101
	// 2 的二进制表示：010
	var (
		m = 5
		n = 2
	)

	// & 按位与
	fmt.Println(m & n)

	// ｜ 按位或
	fmt.Println(m | n)

	// ^ 按位非
	fmt.Println(^m)

	// ^ 按位异或
	fmt.Println(m ^ n)

	// << 按位左移
	fmt.Println(m << 2)
	fmt.Println(1 << 10) // 10 0000 0000 => 1024

	// >> 按位右移
	fmt.Println(m >> 2)
	var p = int8(1)
	fmt.Println(p << 10) // 'p' (8 bits) too small for shift of 10  可以运行

	// 赋值运算符
	var x int
	x = 10
	x += 2
	fmt.Println(x)
	x = 10
	x -= 2
	fmt.Println(x)
	x = 10
	x *= 2
	fmt.Println(x)
	x = 10
	x /= 2
	fmt.Println(x)
	x = 10
	x %= 2
	fmt.Println(x)
	x = 10
	x <<= 2
	fmt.Println(x)
	x = 10
	x >>= 2
	fmt.Println(x)
}
