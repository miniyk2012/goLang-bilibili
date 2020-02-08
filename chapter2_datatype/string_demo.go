package main

import (
	"fmt"
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
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	for _, v := range s {
		fmt.Printf("%v(%c) ", v, v)
	}
	fmt.Println()

}
func main() {
	printSth()
	fmt.Println()
	stringFunPrint()
	fmt.Println()
	bytePrint()
	fmt.Println()
	traversalString("hello world, 你好")
}
