package main

import "fmt"

// goto

func main() {
	breakFlag := false
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				breakFlag = true
				break
			}
			fmt.Println(i, j)
		}
		if breakFlag {
			break
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}

breakTag:
	fmt.Println("break")

breakTag1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break breakTag1
			}
			fmt.Println(i, j)
		}
	}

breakTag2:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue breakTag2
			}
			fmt.Println(i, j)
		}
	}
}
