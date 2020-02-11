// JSON的序列化和反序列化

package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func (c *Class) PrintClass() {
	fmt.Printf("%v\n", c.Title)
	for _, v := range c.Students {
		fmt.Printf("%#v\n", *v)
	}
}

func demo1() {
	var c *Class = new(Class)
	fmt.Printf("%#v, len=%d, cap=%d\n", c.Students, len(c.Students), cap(c.Students))

	c.Title = "三年二班"
	c.Students = make([]*Student, 0, 100) // 可以不写, 似乎new的时候就已经给c.Students分配了空间了
	fmt.Printf("%#v, len=%d, cap=%d\n", c.Students, len(c.Students), cap(c.Students))
	for i := 0; i < 5; i++ {
		c.Students = append(c.Students, &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		})
	}
	fmt.Printf("%#v, len=%d, cap=%d\n", c.Students, len(c.Students), cap(c.Students))
	c.PrintClass()
	println()

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	println()
	var c1 *Class = new(Class) // 需要分配空间哟
	err = json.Unmarshal([]byte(data), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	c1.PrintClass()

	println()
	var c2 Class // c2已经自动分配了空间
	fmt.Printf("%#v\n", c2)
	err = json.Unmarshal([]byte(data), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	c2.PrintClass()

	println()
	c3 := &Class{}
	err = json.Unmarshal([]byte(data), &c3)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	c3.PrintClass()

}

func main() {
	demo1()
}
