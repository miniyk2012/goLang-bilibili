package main

import "fmt"

// 整型、浮点型、布尔型、字符串型变量
func printT() {
	var (
		a string  = "hello"
		b float64 = 100.5
		c bool    = true
		d int     = -103
	)
	fmt.Printf("%s\n", a)
	fmt.Printf("%.2f\n", b)
	fmt.Printf("%t\n", c)
	fmt.Printf("%d\n", d)
}

// 编写代码统计出字符串"hello沙河小王子"中汉字的数量。
func countCnCount(s string) {
	var count = 0
	const startCode = 0x2E80
	const endCode = 0x9FFF
	runeS := []rune(s)
	for _, v := range runeS {
		if v >= startCode && v <= endCode {
			count++
		}
	}
	fmt.Println("中文个数:", count)
}
func main() {
	printT()
	println()
	countCnCount("hello沙河小王子 hi, 你好")
}
