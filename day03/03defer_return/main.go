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
	return x // 1. 返回值赋值 返回值指向 5，没有指向 x 2. defer  3. 真正的RET
}

func f2() (x int) {
	x = 5
	defer func() {
		x++ // 返回值是命名的，返回值和 x 等价
	}()
	return // 1. 返回值赋值 返回值指向x  2. defer x被修改了，相当于返回值被修改了  3. 真正的RET
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ // 返回值 y 赋值为 5，defer 修改的是 x
	}()
	return x // 返回值指向 y，没有指向 x
}

func f4() (x int) {
	x = 5
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return // 返回值指向 x，但修改的是 x 的副本
}

func f5() (x int) {
	defer func(x int) int {
		x++ // 修改的是函数中的副本
		return x
	}(x)
	return 5 // 返回值 = x = 5
}

// 传一个 x 的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 1. 返回值 = x = 5 2. defer x=6 3. RET 返回
}

func main() {
	fmt.Println(f1()) // 5 因为返回值不是命名的，第一步返回值已经复制为 5 了
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
	fmt.Println(f6()) // 6
}
