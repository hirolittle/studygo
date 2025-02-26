package main

import "fmt"

// make 函数创造切片

func main() {

	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 切片不能直接比较，因为是引用类型，只能和 nil 做比较
	// 一个`nil`值的切片并没有底层数组，一个`nil`值的切片的长度和容量都是0
	// 但是我们不能说一个长度和容量都是0的切片一定是`nil`
	// 所以要判断一个切片是否是空的，要是用`len(s) == 0`来判断，不应该使用`s == nil`来判断
	var sli1 []int
	var sli2 = []int{}
	var sli3 = make([]int, 0)
	fmt.Printf("sli1=%v len(sli1)=%d cap(sli1)=%d sli1==nil: %v\n", sli1, len(sli1), cap(sli1), sli1 == nil)
	fmt.Printf("sli2=%v len(sli2)=%d cap(sli2)=%d sli2==nil: %v\n", sli2, len(sli2), cap(sli2), sli2 == nil)
	fmt.Printf("sli3=%v len(sli3)=%d cap(sli3)=%d sli3==nil: %v\n", sli3, len(sli3), cap(sli3), sli3 == nil)

	// 切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 // s3 和 s4 都指向了同一个底层数组
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(i, s3[i])
	}
	for i, v := range s3 {
		fmt.Println(i, v)
	}

}
