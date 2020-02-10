package main

import (
	"fmt"
)

type person struct {
	name, city string
	age        int
}

func newPerson(name, city string, age int) *person {
	return &person{name, city, age} // &main.person{name:"杨恺", city:"常州", age:20}
}

func demo1() {
	p1 := newPerson("杨恺", "常州", 20)
	fmt.Printf("%#v\n", p1)
}

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 方法, 它属于Person
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	// Go语言会在代码运行时将接收者的值复制一份
	p.age = newAge
}

func demo2() {
	p2 := NewPerson("王洋", 40)
	(*p2).Dream()
	p2.Dream() // 语法糖, 效果同上一行
	fmt.Printf("%#v\n", p2)

	p2.SetAge(100)
	fmt.Printf("%#v\n", p2)
	(*p2).SetAge(60) // 能改动哟
	fmt.Printf("%#v\n", p2)

	p2.SetAge2(0)
	fmt.Printf("%#v\n", p2)

	p3 := Person{
		name: "tiancai",
		age:  20,
	}
	p3.SetAge(50) // 也能改动哟, 因此调用方是什么, 看的是方法定义时的接收者是指针还是值
	fmt.Printf("%#v\n", p3)
	(&p3).SetAge(80)
	fmt.Printf("%#v\n", p3)
}

func main() {
	demo1()
	println()
	demo2()
}
