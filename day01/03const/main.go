package main

import "fmt"

// 常量

const pi = 3.1415926

// 批量声明
const (
	statusOK = 200
	notFound = 404
)

// 批量声明时，如果某一行声明后没有赋值，默认和上一行相同
const (
	n1 = 100
	n2 // 100
	n3 // 100
)

// iota
// `iota ` 是 go 语言的常量计数器，只能在常量的表达式中使用
// `iota `在 const 关键字出现时将被重置为0
// const中每新增一行常量声明将使`iota`计数一次
const (
	a = iota
	b // = iota, 1
	c // = iota, 2
	d // = iota, 3
)

// _ 跳过某些值
const (
	b1 = iota
	b2 // = iota, 1
	_  // = iota, 2
	b3 // = iota, 3
)

// 中间插入值
const (
	c1 = iota
	c2 = 100
	c3 = iota // 2
	c4        // = iota, 3
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 // iota == 0, d1 == 1, d2 == 2
	d3, d4 = iota + 1, iota + 2 // iota == 1, d3 == 2, d4 == 3
)

// 定义数量级
const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1024
	MB = 1 << (10 * iota) // 1048576
	GB                    // 1073741824
	TB                    // 1099511627776
	PB                    // 1125899906842624
	EB                    // 1152921504606846976
	//ZB
	//YB
)

func main() {
	// 定义了常量之后，不能修改
	// 在程序运行期间不会改变
	// '1.23' (type untyped float) cannot be represented by the type untyped float
	//pi = 1.23

	fmt.Printf("pi: %f, %T\n", pi, pi)

	fmt.Printf("statusOK: %d, notFound: %d\n", statusOK, notFound)

	fmt.Printf("n1: %d, n2: %d, n3: %d\n", n1, n2, n3)

	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)

	fmt.Printf("b1: %d, b2: %d, b3: %d\n", b1, b2, b3)

	fmt.Printf("c1: %d, c2: %d, c3: %d, c4: %d\n", c1, c2, c3, c4)

	fmt.Printf("d1: %d, d2: %d\n", d1, d2)
	fmt.Printf("d3: %d, d4: %d\n", d3, d4)

	fmt.Printf("KB: %d, MB: %d, GB: %d, TB: %d, PB: %d，EB: %d\n", KB, MB, GB, TB, PB, EB)
}
