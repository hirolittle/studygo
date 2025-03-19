package main

import "fmt"

// 递归：函数自己调用自己！
// 递归适合处理那种问题相同，问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 上台阶
// n 个台阶，一次可以走一步，也可以走两步，有多少种走法

func taijie(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

// 3! = 3 * 2* 1
// 4! = 4 * 3 * 2 * 1
// 5! = 5 * 4 * 3 * 2 * 1 = 5 * 4!

// 计算 n 的阶乘
func f(n uint) uint {
	if n <= 1 {
		return n
	}
	return n * f(n-1)
}

func main() {

	fmt.Println(taijie(3))

	fmt.Println(f(5))

}
