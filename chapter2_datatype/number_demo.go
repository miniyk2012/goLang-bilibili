package main

import (
	"fmt"
	"math"
)

func intPrint() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77
	fmt.Printf("%d \n", b) // 63

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
	fmt.Printf("%d \n", c) // 255

	// 二进制
	var d int64 = 0b1010
	fmt.Printf("%b \n", d) // 1010
	fmt.Printf("%d \n", d) // 10
}

func floatPrint() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}

func complexPrint() {
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
func main() {
	intPrint()
	fmt.Println()
	floatPrint()
	fmt.Println()
	complexPrint()
}
