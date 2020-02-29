package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// main.go

func sayHello1(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello1.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "沙河小王子")
}
func main() {
	http.HandleFunc("/hello1", sayHello1)
	http.HandleFunc("/hello2", sayHello2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello2(w http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./hello2.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := User{"杨恺", "男", 20}
	m := map[string]interface{} {
		"name": "王洋",
		"gender": "女",
		"age": 50,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, map[string]interface{}{
		"u1": user,
		"m1": m,
	})
}