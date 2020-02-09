package main

import (
	"fmt"
)

func pointerDemo() {
	a := 10
	pa := &a
	fmt.Printf("%d, %p\n", a, &a)
	fmt.Printf("%p, type=%T\n", pa, pa)
	fmt.Printf("%p\n", &pa)
	println()
}

func modify(a *int) {
	*a = 100
}

func allocate() {
	var a = new(int)
	fmt.Println(*a)

	var b = make(map[string]int)
	b["沙河娜扎"] = 100
	fmt.Println(b)
	println()

	// 分配内存空间, 设置为0值
	d := new(int)
	e := new(bool)
	fmt.Printf("%T\n", d) // *int
	fmt.Printf("%T\n", e) // *bool
	fmt.Println(*d)       // 0
	fmt.Println(*e)       // false

	// https://www.jianshu.com/p/c173dab0e71c: new和 make的区别: new和 make的区别
	println()

}

func main() {
	pointerDemo()

	a := 10
	println("before", a)
	modify(&a)
	println("after", a)
	println()
	allocate()
}
