package main

import (
	"fmt"
)

type Cat struct{}
type Dog struct {
	name string
}
type Car struct {
	Brand string
}

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

func (dog Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", dog.name)
}

func (car Car) Move() {
	fmt.Printf("%s速度70迈\n", car.Brand)
}

func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.name)
}

func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func (cat Cat) Say() {
	fmt.Println("喵喵!")
}

func da(arg Sayer) {
	arg.Say()
}

func demo1() {
	dog := Dog{}
	cat := Cat{}
	var x Sayer
	x = dog
	da(x)
	x = cat
	da(x)
}

func demo2() {
	dog := Dog{}
	cat := Cat{}
	var x Mover
	x = dog
	x.Move()
	x = &dog
	x.Move()
	// x = cat // 不可以
	x = &cat
	fmt.Printf("type:%T value:%v\n", x, x)
	x.Move()

	cat.Move() // 可以
	dog.Move()

	(&cat).Move() // 可以
	(&dog).Move() // 可以
}

func demo3() {
	var x Sayer
	var y Mover
	dog := Dog{"旺财"}
	x = dog
	x.Say()
	y = dog
	y.Move()
}

func demo4() {
	var x Mover
	c := Car{"奔驰"}
	d := Dog{"小黑"}
	x = c
	x.Move()
	x = d
	x.Move()
}

type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {
	name string
}

type haier struct {
	dryer // 一定要匿名才能继承哟
	name  string
}

func (d dryer) wash() {
	fmt.Println("甩一甩")
}

func (h haier) dry() {
	fmt.Println("洗刷刷")
}

// 一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现
func demo5() {
	var w WashingMachine
	fmt.Printf("type:%T value:%v\n", w, w)
	h := haier{dryer{"干燥"}, "海尔"}
	w = h
	w.dry()
	w.wash()
	fmt.Printf("type:%T value:%v\n", w, w)
}

func main() {
	demo1()
	println()
	demo2()
	println()
	demo3()
	println()
	demo4()
	println()
	demo5()
}
