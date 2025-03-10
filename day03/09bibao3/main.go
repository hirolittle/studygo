package main

import (
	"fmt"
	"strings"
)

// 闭包

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))      //test.jpg
	fmt.Println(jpgFunc("test1.jpg")) //test1.jpg
	fmt.Println(txtFunc("test"))      //test.txt
	fmt.Println(txtFunc("test2.txt")) //test2.txt
}
