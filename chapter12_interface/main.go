package main

import (
	"fmt"
)

type Cat struct{}
type Dog struct{}

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

func (dog Dog) Say() {
	fmt.Println("汪汪!")
}

func (d Dog) Move() {
	fmt.Println("狗会动")
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
	x.Move()

	cat.Move() // 可以
	dog.Move()

	(&cat).Move() // 可以
	(&dog).Move() // 可以
}

func main() {
	demo1()
	println()
	demo2()
}
