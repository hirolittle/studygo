package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map sort key, 按照指定顺序遍历

func main() {
	// 旧方式
	//rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	// 新方式
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成 stu 开头的字符串
		// 旧方式
		//value := rand.Intn(100)             // 生成 0~99 的随机整数
		// 新方式
		value := r.Intn(100) // 生成 0~99 的随机整数
		scoreMap[key] = value
	}

	// 取出 map 中的所有 key
	var keys []string
	for k := range scoreMap {
		keys = append(keys, k)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序后的 key 遍历 map
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, scoreMap[k])
	}

}
