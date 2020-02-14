package main

import (
	"fmt"
	"reflect"
)

func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type:%15v\t", t)
	// v := reflect.ValueOf(a)
	// fmt.Printf("value:%10v\t", v)
	// Kind()指底层类型
	fmt.Printf("Name:%12v\tKind:%15v\n", t.Name(), t.Kind())
}

func demo1() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

type myInt int64

func demo2() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d)
	reflectType(e)
	var f []int
	var g [4]int
	var h map[string]int
	reflectType(f)
	reflectType(g)
	reflectType(h)
	if reflect.TypeOf(f).Kind() == reflect.Slice {
		fmt.Println("f is a Slice.")
	}

}

func main() {
	demo1()
	println()
	demo2()
}
