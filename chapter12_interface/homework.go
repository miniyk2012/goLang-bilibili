package main

import (
	"fmt"
	"io"
	"os"
)

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
	//var peo People = Student{} // 不可以
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。
type Logger interface {
	Info(string)
}

type FileLogger struct {
	filename string
}

func (fl *FileLogger) Info(msg string) {
	var f *os.File
	var err1 error
	if checkFileIsExist(fl.filename) { //如果文件存在
		f, err1 = os.OpenFile(fl.filename, os.O_APPEND|os.O_WRONLY, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(fl.filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	n, err1 := io.WriteString(f, msg+"\n") //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节\n", n)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) Info(msg string) {
	fmt.Println(msg)
}

func homework2() {
	var logger Logger
	fileLogger := &FileLogger{"log.txt"}
	logger = fileLogger
	logger.Info("Hello")
	logger.Info("how are you")

	consoleLogger := &ConsoleLogger{}
	logger = consoleLogger
	logger.Info("Hello")
	logger.Info("how are you")
}

func main() {
	homework1()
	println()
	homework2()
}
