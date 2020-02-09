package main

import (
	"fmt"
	"strings"
)

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

func rules(name string) (num int) {
	countChar := func(char string) (count int) {
		for _, c := range name {
			fmt.Println(string(c))
			if char == string(c) {
				count++
			}
		}
		return count
	}
	num += countChar("e") + countChar("E")
	num += 2 *  (countChar("i") + countChar("I"))
	num += 3 * (countChar("o") + countChar("O"))
	num += 4 * (countChar("u") + countChar("U"))

	return
}

func dispatchCoin() int {
	var distributed int
	for _, v := range users {
		distributed += rules(v)
		distribution[v] = rules(v)
	}
	fmt.Println(distribution)
	return coins - distributed
}

func dispatchCoin2() int {
	left := coins
	for _, name := range users {
		e := strings.Count(name, "e") + strings.Count(name, "E")
		i := strings.Count(name, "i") + strings.Count(name, "I")
		o := strings.Count(name, "o") + strings.Count(name, "O")
		u := strings.Count(name, "u") + strings.Count(name, "U")
		sum := e*1 + i*2 + o*3 + u*4
		distribution[name] = sum
		left -= sum
	}
	return left
}
func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)

	left = dispatchCoin2()
	fmt.Println("剩下：", left)
}
