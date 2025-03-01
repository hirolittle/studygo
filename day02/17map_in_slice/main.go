package main

import "fmt"

// map in slice
// 元素为map类型的切片

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}
	fmt.Println("after init")

	// 对切片中的 map 元素进行初始化
	mapSlice[0] = make(map[string]string, 3)
	mapSlice[0]["name"] = "hiro"
	mapSlice[0]["age"] = "20"
	mapSlice[0]["address"] = "beijing"

	for index, value := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}
}
