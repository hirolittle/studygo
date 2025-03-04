package main

import "fmt"

// defer 经典案例

// Go 语言中函数的 return 不是原子操作，在底层是分为两步来执行
// 第一步：返回值赋值
// 第二步：真正的 RET 返回
// 函数中如果存在 defer，那么 defer 执行的时机在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 返回值不是命名的，defer 修改的是x，不是返回值
	}()
	return x
}

func f2() (x int) {
	x = 5
	defer func() {
		x++ // 返回值是命名的，返回值和 x 等价
	}()
	return
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 返回值 y 赋值为 5，defer 修改的是 x
	}()
	return x
}
func f4() (x int) {
	x = 5
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return // 返回值 == x == 5
}
func main() {
	fmt.Println(f1()) // 因为返回值不是命名的，第一步返回值已经复制为 5 了
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
