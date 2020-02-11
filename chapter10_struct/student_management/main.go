package main

import (
	"fmt"
	"os"
)

func showMenu() int {
	fmt.Println("学生信息管理系统, 请输入操作类型:")
	fmt.Println("1: 增加学生")
	fmt.Println("2: 修改学生")
	fmt.Println("3: 删除学生")
	fmt.Println("4: 显示学生")
	fmt.Println("5: 退出")
	var c = new(int)
	fmt.Scanf("%d", c)
	return *c
}

func prompt() *Student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("id: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("class: ")
	fmt.Scanf("%s\n", &class)
	return &Student{id, name, class}
}

func main() {
	smg := newStudentManagement()
	for {
		input := showMenu()
		switch input {
		case 1:
			// 增加学生
			stu := prompt()
			smg.addStudent(stu)
		case 2:
			// 修改学生
			stu := prompt()
			smg.modifyStudent(stu)
		case 3:
			// 删除学生
			var id int
			fmt.Print("id: ")
			fmt.Scanf("%d\n", &id)
			smg.delStudent(id)
		case 4:
			// 显示学生
			smg.showStudents()
		case 5:
			os.Exit(0)
		}
	}
}
