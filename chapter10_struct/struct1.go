package main

import (
	"fmt"
	"unsafe"
)

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

func typeAlias() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int, MyInt只会在代码中存在，编译完成时并不会有MyInt类型
}

type Person struct {
	name, city string
	age        int
}

func usePerson1() {
	a := Person{} // 分配空间
	b := &a
	fmt.Printf("%#v, %#v\n", a, *b)
	a = Person{"yang", "changzhou", 10} // a是引用, 绑定的对象不会再变
	fmt.Printf("%#v, %#v\n", a, *b)
	*b = Person{"yangkai", "changzhou", 10}
	fmt.Printf("%#v, %#v\n", a, *b)
}

func usePerson2() {
	var p1 Person
	fmt.Printf("p1=%v\n", p1)
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
}

func anonymousStru() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "yang"
	user.Age = 13
	fmt.Printf("%v\n", user)
}

func newStruct() {
	p2 := new(Person)
	fmt.Println(p2)
	p2.name = "yay"
	p2.age = 12
	fmt.Printf("%#v\n", *p2)
	println()

	p3 := &Person{} // 也会分配内存
	fmt.Printf("%#v\n", p3)
	p3.name = "杨恺"
	p3.age = 12
	p3.city = "常州"
	(*p3).city = "上海"
	fmt.Printf("%#v\n", p3)

	var p4 Person
	fmt.Printf("%#v\n", p4)

	var p5 = Person{city: "changzhou", age: 12, name: "养成"}
	fmt.Printf("%#v\n", p5)

	var p6 = &Person{city: "changzhou", age: 123, name: "养成a "}
	fmt.Printf("%#v\n", p6)

	var p7 = &Person{city: "changzhou市"}
	fmt.Printf("%#v\n", p7)

	var p8 = &Person{"杨恺", "常州", 13}
	fmt.Printf("%#v\n", p8)

}

type test struct {
	a, b, c, d int8
}

// 内存对齐
func memoryTest() {
	t := test{1, 2, 3, 4}
	fmt.Printf("%p\n", &t.a)
	fmt.Printf("%p\n", &t.b)
	fmt.Printf("%p\n", &t.c)
	fmt.Printf("%p\n", &t.d)

	println()
	fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	var s string = "ABDCD"
	for i := 0; i < 10; i++ {
		s = fmt.Sprintf("%v%v", s, s)
	}
	fmt.Printf("string size: %d\n", unsafe.Sizeof(s))
}

func main() {
	typeAlias()
	println()
	usePerson1()
	println()
	usePerson2()
	println()
	anonymousStru()
	println()
	newStruct()
	println()
	memoryTest()
}
