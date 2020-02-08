package main

import (
	"fmt"
)

var (
	m         = 100
	name, age = "Q1mi", 20
	a         string
	b         int
	c         bool
	d         float32
)

func foo() (int, string) {
	return 10, "Q1"
}

func main() {
	n := 100
	m := 200
	fmt.Println(m, n)
	fmt.Println("name=", name)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
}
