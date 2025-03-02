package main

import "fmt"

// 判断回文字符串
// 字符串从左往右读和从右往左读是一样的，那么就是回文
// 上海自来水来自海上

func main() {
	s := "a上海自来水来自海上a"
	var sli []rune
	for _, c := range s {
		sli = append(sli, c)
	}

	// 1. 遍历一半字符
	for i := 0; i < len(sli)/2+1; i++ {
		// 2. 第一个字符和最后一个相等，第n个字符和len(s)-1个相等
		if sli[i] != sli[len(sli)-1-i] {
			fmt.Println("不是回文字符串")
			return
		}
	}

	fmt.Println("是回文字符串")

}
