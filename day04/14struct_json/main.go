package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与 JSON

// 1. 把 Go 中的结构体变量 --> json 格式的字符串
// 2. json 格式的字符串 --> Go 语言中能够识别的结构体

// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）

type person struct {
	Name string `json:"name"` // 通过指定tag实现json序列化该字段时的key
	Age  int    `json:"age"`  //json序列化是默认使用字段名作为key
	//gender string `json:"gender"` //私有不能被json包访问
}

func main() {

	p1 := person{
		Name: "John Doe",
		Age:  42,
	}

	// json 序列化
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	// json 反序列化
	s := `{"name": "hiro", "age": 18}`
	var p2 person
	// 传指针，是为了能在 json.unmarshal 内部修改 p2 的值
	err = json.Unmarshal([]byte(s), &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
