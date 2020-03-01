package main

import "fmt"

func enemy() func() {
	key := 0
	return func() {
		key++
		fmt.Println(key)
	}
}

func adder(a, b int) int {
	return a + b
}


func main() {
	f := enemy()
	f()
	f()
	f()
	f()

	fiveAdd := curryAdder(5)
	fmt.Println(fiveAdd(10))
	fmt.Println(fiveAdd(15))
}

func curryAdder(a int) func (b int) int {
	return func(b int) int {
		return adder(a, b)
	}
}