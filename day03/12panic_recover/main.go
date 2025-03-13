package main

import "fmt"

// panic and recover

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("func B recover:", err)
		}
	}()
	panic("func B panic") // 程序崩溃退出
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
