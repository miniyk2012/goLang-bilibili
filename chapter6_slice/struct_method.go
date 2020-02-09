package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	sam.SayHi()
	sam.Human.SayHi()    // calls Human.SayHi
	sam.SayHello()       // calls  Human.SayHello
	sam.Human.SayHello() // calls  Human.SayHello
}
