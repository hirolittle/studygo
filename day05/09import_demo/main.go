package main

import (
	"fmt"
	"github.com/hirolittle/studygo/day05/08calc"
)

// go modules 引入本地包
// go mod edit -replace github.com/hirolittle/studygo/day05/08calc=../08calc
// go mod tidy

func main() {
	res := calc.Add(1, 2)
	fmt.Println(res)
}
