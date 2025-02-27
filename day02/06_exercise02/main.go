package main

import "fmt"

// 找出数组中和为指定值的两个元素的下标，比如从数组 [1, 3, 5, 7, 8] 中找出和为 8 的两个元素的下标为 (0, 3) 和 (2, 4)

func main() {
	arr := [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}

	res := make(map[int]int)
	target := 8
	for i, num := range arr {
		item := target - num
		if j, ok := res[item]; ok {
			fmt.Printf("(%d, %d)\n", j, i)
		}
		res[num] = i
	}

}
