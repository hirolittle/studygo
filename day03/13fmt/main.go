package main

import "fmt"

// fmt

func printBaifenbi(num int) {
	fmt.Printf("%d%%\n", num)
}

func main() {
	fmt.Print("沙河")
	fmt.Print("娜扎")
	fmt.Println()

	fmt.Println("沙河")
	fmt.Println("娜扎")

	// Printf("格式化字符串", 值)
	// %T: 查看类型
	// %d: 十进制数
	// %b: 二进制数
	// %o: 八进制数
	// %x: 十六进制数
	// %c: 字符
	// %s: 字符串
	// %p: 指针
	// %v: 值
	// %f: 浮点数
	// %t: 布尔值
	var m1 = make(map[string]int, 1)
	m1["hiro"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%#v\n", m1)

	printBaifenbi(90)

	fmt.Printf("%t\n", true)

	// 整数 -> 字符
	// 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", 65)

	// 宽度标识符
	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.f\n", n)
	fmt.Printf("%9.2f\n", n)

	s := "李想有理想"
	fmt.Printf("%q\n", s)
	fmt.Printf("%2.1s\n", s)

	// 获取用户输入
	//var s1 string
	//fmt.Scan(&s1)
	//fmt.Println("input: ", s1)

	var (
		name  string
		age   int
		class string
	)
	//fmt.Scanf("%s %d %s", &name, &age, &class)
	//fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
