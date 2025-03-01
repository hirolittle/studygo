package main

import (
	"fmt"
	"strings"
)

// 写一个程序，统计一个字符串中每个单词出现的次数，比如：“how do you do” 中 how=1, do=2, you=1

func main() {
	s := "how do you do"
	items := strings.Split(s, " ")
	m := make(map[string]int)
	for _, item := range items {
		m[item]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
