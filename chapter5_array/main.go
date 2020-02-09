package main

import (
	"fmt"
	"reflect"
)

func arrayDemo() {
	var a [3]int
	var b [4]int
	c := [3]int{1, 2, 3}
	d := [...]int{5, 6, 7}
	var e = [...]int{1: 4, 3: 5} // [0, 4, 0, 5]
	var f = [...]bool{5: true}   // [false false false false false true]
	fmt.Println(a)
	fmt.Println(b)
	a = c
	fmt.Println(a)
	a = d
	fmt.Println(a)
	fmt.Println(e)
	fmt.Println(f)
}

func twoDimArray() {
	var array = [...][3]int{
		{1, 2, 3}, // 因为外层数组定义时, 已知是[3]int了, 因此可以省略类型
		[...]int{4, 5, 6},
		[3]int{7, 8, 9},
	}
	fmt.Println(array)
	fmt.Println(array[0][2])

	for _, line := range array {
		for _, num := range line {
			fmt.Printf("%d\t", num)
		}
		println()
	}
}

func modify(a [3]int) {
	a[2] = 4
	fmt.Println(a)
}

func equalDemo() {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [...]int{1, 2, 5}
	d := [...]int{1, 2}
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(reflect.DeepEqual(a, d))
	fmt.Println(reflect.DeepEqual(a, b))

}

func assignArray() {
	a := [...]int{1, 2, 3}
	b := a // 复制赋值
	a[0] = 10
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	arrayDemo()
	println()
	twoDimArray()
	println()
	a := [3]int{1, 2, 3}
	modify(a)
	fmt.Println(a)
	println()
	equalDemo()
	println()
	assignArray()
}
