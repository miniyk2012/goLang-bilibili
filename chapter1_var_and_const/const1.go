package main

import (
	"fmt"
)

const (
	pi      = 3.1415
	epsilon = 2.7182
)

const (
	n1 = 100
	n2
	n3
)

// itoa起始值为0, 每新起一行+1
const (
	a1 = iota // 0
	a2 = 100  // 100
	a3 = iota // 2
	_
	a4 // 4
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 2 ** 10
	MB = 1 << (10 * iota) // 2 ** 20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func constPrint() {
	fmt.Println("n1 =", n1)
	fmt.Println("n2 =", n2)
	fmt.Println("n2 =", n2)

}
func iotaPrint() {
	fmt.Println("a1 =", a1)
	fmt.Println("a2 =", a2)
	fmt.Println("a3 =", a3)
	fmt.Println("a4 =", a4)
}

func storagePrint() {
	fmt.Println("KB =", KB)
	fmt.Println("MG =", MB)
	fmt.Println("GB =", GB)
	fmt.Println("TB =", TB)
	fmt.Println("PB =", PB)
}

func doublePrint() {
	fmt.Println("a =", a, ",b =", b)
	fmt.Println("c =", c, ",d =", d)
	fmt.Println("e =", e, ",f =", f)
}

func main() {
	constPrint()
	fmt.Println()
	iotaPrint()
	fmt.Println()
	storagePrint()
	fmt.Println()
	doublePrint()
}
