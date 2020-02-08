package main

import (
	"fmt"
	"math"
	"strings"
)

func printSth() {
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	var a string = `牛逼
红红`
	fmt.Println(a)
}
func stringFunPrint() {
	var a string = "hello world"
	fmt.Println(len(a))
	fmt.Println(strings.Split(a, " "))
	fmt.Println(strings.Contains(a, "he"))
	fmt.Println(strings.Join([]string{"you", "are", "good"}, " "))
}

func bytePrint() {
	var a rune = '中'
	var b byte = 'x'
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", b)
}

func traversalString(s string) {
	fmt.Println("len(s) =", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	for _, v := range s {
		fmt.Printf("%v(%c) ", v, v)
	}
	fmt.Println()
}

func changeString(s string) {
	byteS1 := []byte(s)
	byteS1[0] = 'w'
	fmt.Println(string(byteS1))

	runeS := []rune(s)
	runeS[0] = '哦'
	fmt.Println(string(runeS))
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println("c = ", c)
}

func main() {
	printSth()
	fmt.Println()
	stringFunPrint()
	fmt.Println()
	bytePrint()
	fmt.Println()
	traversalString("hi, 你好")
	println()
	changeString("hi, 你好")
	changeString("喂, 你好")
	println()
	sqrtDemo()
}
