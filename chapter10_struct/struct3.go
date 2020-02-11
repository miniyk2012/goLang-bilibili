package main

import (
	"fmt"
)

type Address struct {
	Province   string
	City       string
	CreateTime string
}

func (a *Address) sayHello() {
	fmt.Printf("Hello, I'm %s\n", a.City)
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

type User struct {
	Name   string
	Gender string
	Address
	Email
}

func demo1() {
	user := User{
		Name:   "杨恺",
		Gender: "男",
		Address: Address{
			City:     "常州",
			Province: "江苏",
		},
	}
	fmt.Printf("%#v\n", user)
	fmt.Printf("%s\n", user.City) // 继承了匿名结构体的字段
	user.sayHello()               // 继承了匿名结构体的方法
	println()

	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东" //通过匿名结构体.字段名访问
	user2.City = "威海"             //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2)

	println()
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019"         //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "3000"   //指定Email结构体中的CreateTime
	fmt.Printf("user2=%#v\n", user3)
}

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会跑~\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name) // 可以直接访问匿名结构中的字段
}

// 实现继承
func demo2() {
	dog := &Dog{
		Feet:   10,
		Animal: &Animal{"妞妞"},
	}
	dog.wang()
	dog.move() // 可以直接使用匿名结构的方法
	dog.Animal.move()
}

func main() {
	demo1()
	println()
	demo2()
}
