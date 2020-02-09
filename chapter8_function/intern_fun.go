package main

import (
	"fmt"
)

func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("%T, %v\n", err, err)
		}
	}()
	panic("panic b")
}

func c() {
	fmt.Println("func c")
}

// https://www.jianshu.com/p/c173dab0e71c: new和 make的区别
func newDemo() {
	p := new([]int) // 分配一块空间存储slice的(ptr, len, cap)三元组(都初始化为零值), 然后p指向这块空间(可以称作指向slice的指针)
	fmt.Println(p)
	fmt.Println(*p, len(*p), cap(*p))

	v := make([]int, 0) // v is initialed with len 10, cap 50
	fmt.Println(v, len(v), cap(v))
	println()
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", *p)
	fmt.Printf("%T\n", v)
	println()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		*p = append(*p, 18, 19, 20)
		fmt.Println(*p, len(*p), cap(*p))
	}()
	(*p)[0] = 18 // panic

}

func main() {
	a()
	b()
	c()
	println()
	newDemo()
}