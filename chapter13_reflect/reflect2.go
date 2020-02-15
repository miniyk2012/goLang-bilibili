package main

import (
	"reflect"
	"fmt"
)

type student struct {
	Name  string `json:"name" ini:"value"`
	Score int    `json:"score" ini:"xx"`
	value int
}

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func (s  student) Eat(food string) {
	fmt.Printf("eat %s\n", food)
}

func demo1() {
	var s  student
	t := reflect.TypeOf(s)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i:=0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name: %s, PkgPath: %s, Type: %v, Index: %d, tag json: %v\n", 
		field.Name, field.PkgPath, field.Type, field.Index, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v, ini tag:%v\n",
			scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"), scoreField.Tag.Get("ini"))
	}
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())
	fmt.Println(v.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("%s, %s, %s\n", t.Method(i).Name, t.Method(i).Type, v.Method(i).Type())
	}
	// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
	var arg = []reflect.Value{}
	v.Method(1).Call(arg)

	v.MethodByName("Study").Call(nil)
	eatFunc := v.MethodByName("Eat")
	eatFunc.Call([]reflect.Value{reflect.ValueOf("狗屎")})
}

func demo2() {
	stu := student{
		Name:  "小王子",
		Score: 90,
	}
	printMethod(stu)
}

func main() {
	demo1()
	println()
	demo2()
}