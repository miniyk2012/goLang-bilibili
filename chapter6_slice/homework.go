package main

import (
	"fmt"
)

func h1(i int) {
	var a = make([]string, i, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%v, len=%d, cap=%d\n", a, len(a), cap(a))
}

func main() {
	h1(5)
	h1(0)
	println()
	var a = [...]int{3, 7, 8, 9, 1}
	fmt.Println(a)
}
