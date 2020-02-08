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

func main() {
	printSth()
	fmt.Println()
	stringFunPrint()
}
