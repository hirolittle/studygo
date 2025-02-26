package main

import "fmt"

// switch
// 简化大量的判断

func main() {

	// if 判断
	var finger = 5
	if finger == 1 {
		fmt.Println("大拇指")
	} else if finger == 2 {
		fmt.Println("食指")
	} else if finger == 3 {
		fmt.Println("中指")
	} else if finger == 4 {
		fmt.Println("无名指")
	} else if finger == 5 {
		fmt.Println("小拇指")
	} else {
		fmt.Println("无效的数字")
	}

	// switch 简化上面的代码
	var finger1 = 5
	switch finger1 {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	switch finger2 := 5; finger2 {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	// 一个分支可以有多个值，多个 case 值中间使用英文逗号分隔
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	// 分支还可以使用表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// fallthrough 语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
