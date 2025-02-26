package main

import "fmt"

// 求数组 [1, 3, 5, 7, 8] 所有元素的和

func main() {

	arr := []int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)

}
