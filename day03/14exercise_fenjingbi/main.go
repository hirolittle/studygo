package main

import "fmt"

// 分金币

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
	a. 名字中每包含1个'e'或'E'分1枚金币
	b. 名字中每包含1个'i'或'I'分2枚金币
	c. 名字中每包含1个'o'或'O'分3枚金币
	d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	// 1. 获取每一个人名字
	for _, name := range users {
		// 2. 获取名字里面每个字符
		for _, c := range name {
			// 3. 匹配字符，分配金币
			switch c {
			case 'e', 'E':
				distribution[name] += 1
				coins -= 1
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			default:
				continue
			}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下", left)

	for k, v := range distribution {
		fmt.Println(k, v)
	}
}
