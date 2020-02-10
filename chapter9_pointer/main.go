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
	var pb = &b
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

	modifyMap(b)
	fmt.Printf("%p, %v\n", pb, *pb)
	fmt.Println(b)
}

func modifyMap(m map[string]int) {
	m["杨恺"] = 50
	fmt.Printf("%p, %v\n", &m, m)
}

var ap *int

/*
https://spf13.com/post/go-pointers-vs-references/
A pointer is a variable which stores the address of another variable.
A reference is a variable which refers to another variable.
Pointers can be reassigned while references cannot. In other words, a pointer can be assigned to a different address.
*/
func pointerReference() {
	a := 1 // define int
	b := 2 // define int

	ap = &a
	// set ap to address of a (&a)
	//   ap address: 0x2101f1018
	//   ap value  : 1
	fmt.Printf("%p, %v, %v\n", ap, *ap, a)

	*ap = 3
	// change the value at address &a to 3
	//   ap address: 0x2101f1018
	//   ap value  : 3
	fmt.Printf("%p, %v, %v\n", ap, *ap, a)

	a = 4 // 说明a是引用, 因为a改变后, *ap也改变了(引用绑定对象后, 不会再绑定到新的对象上)
	// change the value of a to 4
	//   ap address: 0x2101f1018
	//   ap value  : 4
	fmt.Printf("%p, %v, %v\n", ap, *ap, a)

	ap = &b
	// set ap to the address of b (&b)
	//   ap address: 0x2101f1020
	//   ap value  : 2
	fmt.Printf("%p, %v, %v, %v\n", ap, *ap, a, b)
	println()
	ap2 := ap
	// set ap2 to the address in ap
	//   ap  address: 0x2101f1020
	//   ap  value  : 2
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 2
	fmt.Printf("%p, %v, %p, %v\n", ap, *ap, ap2, *ap2)

	*ap = 5
	// change the value at the address &b to 5
	//   ap  address: 0x2101f1020
	//   ap  value  : 5
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 5
	// If this was a reference ap & ap2 would now
	// have different values
	fmt.Printf("%p, %v, %p, %v\n", ap, *ap, ap2, *ap2)

	ap = &a
	// change ap to address of a (&a)
	//   ap  address: 0x2101f1018
	//   ap  value  : 4
	//   ap2 address: 0x2101f1020
	//   ap2 value  : 5
	// Since we've changed the address of ap, it now
	// has a different value then ap2
	fmt.Printf("%p, %v, %p, %v\n", ap, *ap, ap2, *ap2)

	c := b // c虽然是个引用, 但b被复制了一个新的值给c, 因此c不再和b指向同一个对象
	*ap2 = 10
	fmt.Printf("b=%d, c=%d\n", b, c)

	println()
	s := "ABC"
	ps := &s
	fmt.Printf("%v\n", *ps) // ABC
	s = "CDE"
	fmt.Printf("%v\n", *ps) // CDE
	o := s
	*ps = "EFG"
	fmt.Printf("%v\n", *ps) // EFG
	fmt.Printf("%v\n", s)   // EFG
	fmt.Printf("%v\n", o)   // CDE
}

func main() {
	pointerDemo()

	a := 10
	println("before", a)
	modify(&a)
	println("after", a)
	println()
	allocate()
	println()
	pointerReference()
}
