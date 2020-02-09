package main

import (
	"fmt"
)

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func cal(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
	num--
}

func testLocal() {
	if x := 10; x > 0 {
		z := 5
		fmt.Println(z)
	}
	// fmt.Println(z)//此处无法使用变量z
}

type calulate func(x, y int) int

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

func funcTypeDemo() {
	var c calulate = add
	fmt.Printf("c is %T\n", c)
	fmt.Printf("add is %T\n", add)
	fmt.Println(c(10, 34))
	f := add
	fmt.Printf("f is %T\n", f)
	fmt.Println(f(10, 34))

}

func main() {
	ret1 := intSum3(10)
	ret2 := intSum3(10, 20, 30)
	ret3 := intSum3(10, []int{1, 2, 3}...)
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(ret3)
	println()

	sum, sub := cal(10, 30)
	fmt.Println(sum, sub)
	println()

	testGlobalVar()             // num=10
	fmt.Printf("num=%d\n", num) // num=9
	println()

	testLocal()
	println()

	funcTypeDemo()
}
