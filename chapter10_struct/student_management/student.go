package main

import "fmt"

type Student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *Student {
	return &Student{id, name, class}
}

type StudentManagement struct {
	students []*Student
}

func newStudentManagement() *StudentManagement {
	return &StudentManagement{
		make([]*Student, 0, 100),
	}
}
func (s *StudentManagement) addStudent(student *Student) {
	for _, v := range s.students {
		if v.id == student.id {
			fmt.Printf("id=%d的学生已存在", student.id)
			return
		}
	}
	s.students = append(s.students, student)
}

func (s *StudentManagement) modifyStudent(student *Student) {
	for i, v := range s.students {
		if v.id == student.id {
			s.students[i] = student
			return
		}
	}
	fmt.Printf("未搜索到id=%d学生", student.id)
}

func (s *StudentManagement) delStudent(id int) {
	var (
		index int
		v     *Student
	)
	for index, v = range s.students {
		if v.id == id {
			s.students = append(s.students[:index], s.students[index+1:]...)
			return
		}
	}
	fmt.Printf("未搜索到id=%d学生", id)
}

func (s *StudentManagement) showStudents() {
	for _, stu := range s.students {
		fmt.Printf("id=%d, name=%s, class=%s\n", stu.id, stu.name, stu.class)
	}
}
