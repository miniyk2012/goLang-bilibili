package main

import (
	"reflect"
	"fmt"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
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
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func main() {
	demo1()
}