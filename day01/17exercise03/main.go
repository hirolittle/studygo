package main

import "fmt"

// 99 乘法表

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%d * %d = %d    ", i, j, i*j)
		}
		fmt.Println()
	}
}
