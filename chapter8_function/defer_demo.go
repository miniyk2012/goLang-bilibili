package main

import "fmt"

/*
	在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
	也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行.
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc2(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 很有意思, calc2的第三个参数的calc2(...)是在return前算的
func magic() {
	x := 1
	y := 2
	defer calc2("AA", x, calc2("A", x, y))
	x = 10
	defer calc2("BB", x, calc2("B", x, y))
	y = 20
	fmt.Println("return")
}

func main() {
	deferDemo()
	println()

	// https://www.jianshu.com/p/b4fb3d361d87 见本文的分析
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5

	println()
	magic()
}
