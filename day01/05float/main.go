package main

import (
	"fmt"
	"math"
)

// 浮点数

func main() {
	// float32 最大值，是个 const 常量
	fmt.Printf("%.2f, %g, %.2e\n", math.MaxFloat32, math.MaxFloat32, math.MaxFloat32)

	f1 := 1.23456789 // Go 语言默认小数都是 float64 类型
	fmt.Printf("%f, %T\n", f1, f1)

	// 定义 float32 类型的值
	f2 := float32(1.23456789)
	fmt.Printf("%f, %T\n", f2, f2)

	// float32 类型的值，不能直接赋值给 float64 类型
	// Cannot use 'f2' (type float32) as the type float64
	//f1 = f2

}
