package main

import "fmt"

type person struct {
	name string
}

func (p person)sing() {
	fmt.Println("sing")
}

func (p person)dance() {
	fmt.Println("dance")
}

func (p person)rap() {
	fmt.Println("rap")
}

func (p person)playBasketBall() {
	fmt.Println("playBasketBall")
}

func (p person)fly() {
	fmt.Println("fly")
}

func show(t Trainee) {
	t.sing()
	t.dance()
	t.rap()
	t.playBasketBall()
}

func wow(s Superer){
	s.sing()
	s.fly()
}

type Trainee interface {
	sing()
	dance()
	rap()
	playBasketBall()
}

type Flyer interface {
	fly()
}

type Superer interface {
	Flyer
	Trainee
}


func main() {
	var p person
	show(p)
	println()
	wow(p)
}
