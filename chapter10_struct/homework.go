package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func saveStudents() {
	m := make(map[string]*student)
	students := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, v := range students {
		// go语言坑之for range: https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html
		// 根本原因在于for-range会使用同一块内存去接收循环中的值。
		// 引用v一开始绑定到一个对象后, 后续只是对该对象的值做更新.
		// 因此v的地址永远是不变的
		fmt.Printf("%p\n", &v)
		m[v.name] = &v
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	println("\n等效于:")
	// 等效于
	var value student
	for _, value = range students {
		fmt.Printf("%p\n", &value)
		m[value.name] = &value
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	println()
	fmt.Printf("%p\n", &value)
	fmt.Printf("%p\n", &students[0])
	fmt.Printf("%p\n", &students[1])
	fmt.Printf("%p\n", &students[2])

	println()
	for _, v := range students {
		new_v := v
		fmt.Printf("%p\n", &new_v)
		m[v.name] = &new_v
	}
	students[0].name = "杨恺"
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	fmt.Println(students)

	println()
	students2 := []*student{
		&student{name: "小王子", age: 18},
		&student{name: "娜扎", age: 23},
		&student{name: "大王八", age: 9000},
	}
	for _, v := range students2 {
		m[v.name] = v
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	println()
	students2[0].name = "杨恺"
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func main() {
	saveStudents()
}
