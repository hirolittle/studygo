package main

import "fmt"

// append 为切片追加元素

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	//s1[3] = "广州" // 错误的写法，会导致编译错误：索引越界  panic: runtime error: index out of range [3] with length 3

	// 调用 append 函数必须用原来的切片变量接收返回值
	// append 追加元素，原来的底层数组放不下的时候，Go 底层就会把底层数组换一个
	// 必须用变量接收 append 的返回值
	s1 = append(s1, "广州")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "杭州", "成都")
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	ss := []string{"武汉", "西安", "苏州"}
	s1 = append(s1, ss...) // ... 表示拆开
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

}
