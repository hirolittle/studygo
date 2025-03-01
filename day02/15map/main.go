package main

import "fmt"

// map

func main() {

	// make 初始化
	var m1 map[string]int
	fmt.Println(m1 == nil)        // true, 还有初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) // 要估算好 map 容量，避免在程序运行期间再动态扩容
	m1["hiro"] = 18
	m1["tom"] = 20
	m1["jerry"] = 19
	fmt.Println(m1)
	fmt.Println(m1["hiro"])

	// 在声明的时候填充元素
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 判断某个键是否存在
	fmt.Println(m1["hiro1"]) // 0，如果不存在这个 key，拿到对应类型的零值
	v1, ok := m1["hiro1"]
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(v1)
	}

	// 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 只遍历 key
	for k := range m1 {
		fmt.Println(k)
	}

	// 只遍历 value
	for _, v := range m1 {
		fmt.Println(v)
	}

	// 删除
	delete(m1, "hiro")
	fmt.Println(m1)
	delete(m1, "hiro1") // 删除不存在的 key，不报错

}
