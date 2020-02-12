package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func homework1() {
	var peo People = Student{} // 不可以
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func main() {
	homework1()
}
