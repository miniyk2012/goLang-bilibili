package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

// 接口嵌套
func demo1() {
	var a animal
	c := cat{}
	a = c
	a.move()
	a.say()

	var s Sayer
	s = c
	s.say()

	var m Mover
	m = c
	m.move()
}

// 空接口
func demo2() {
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
	x = &s
	fmt.Printf("type:%T value:%v\n", x, x)
	x = &i
	fmt.Printf("type:%T value:%v\n", x, x)
	x = &b
	fmt.Printf("type:%T value:%v\n", x, x)

	println()
	m := make(map[string]interface{})
	m["a"] = 100
	m["b"] = true
	m["c"] = []string{"a", "b"}
	m["d"] = "海洋可"
	fmt.Println(m)
}

// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
func demo3() {
	var w io.Writer
	w = os.Stdout
	fmt.Printf("type:%T value:%v\n", w, w)
	w = nil
	fmt.Printf("type:%T value:%v\n", w, w)
	w = &bytes.Buffer{}
	fmt.Printf("type:%T value:%v\n", w, w)

	println()
	if v, ok := w.(*bytes.Buffer); ok {
		fmt.Printf("%T\n", v)
	} else {
		fmt.Println("类型断言失败")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var x interface{}
	x = "Hello 沙河"
	fmt.Printf("%v\n", x.(string))

	s := x.(string)
	fmt.Println(s)

	fmt.Printf("%v\n", x.(int))
}

func justifyType(v interface{}) {
	switch x :=v.(type) {
	case string:
		fmt.Printf("v is a string，value is %v\n", x)
	case int:
		fmt.Printf("v is a int is %v\n", v)
	case bool:
		fmt.Printf("v is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")

	}
}

func demo4() {
	justifyType("a")
	justifyType(true)
	justifyType(100)
	justifyType(100.4)
}

func main() {
	demo1()
	println()
	demo2()
	println()
	demo3()
	println()
	demo4()
}
