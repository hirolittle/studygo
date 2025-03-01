package main

import "fmt"

// slice in map
// 值为切片类型的 map

func main() {

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")

	key := "hiro"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0)
	}
	value = append(value, "beijing")
	value = append(value, "shanghai")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
