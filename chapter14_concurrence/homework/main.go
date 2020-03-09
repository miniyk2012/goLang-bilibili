package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
)

const NUM = 100
const WORKER = 5

type Data struct {
	num     int64
	total   int64
	routine int
}

func (d Data) String() string {
	return fmt.Sprintf("sum(%d)=%d, worker=%d", d.num, d.total, d.routine)
}

func calNumSum(num int64) int64 {
	var ret int64
	for num > 0 {
		ret += num % 10
		num = num / 10
	}
	return ret
}

func homework1() {
	var wg1 sync.WaitGroup
	jobs := make(chan int64)

	go func() {
		for i := 0; i < NUM; i++ {
			jobs <- rand.Int63n(1000)
		}
		close(jobs)
	}()

	results := make(chan Data)
	wg1.Add(WORKER)
	go func() {
		wg1.Wait()
		close(results)
	}()

	for i := 0; i < WORKER; i++ {
		go func(i int) {
			defer wg1.Done()
			for num := range jobs {
				total := calNumSum(num)
				results <- Data{num, total, i}
			}
		}(i)
	}

	var line int
	for data := range results {
		fmt.Printf("%d: %v\n", line, data)
		line++
	}
}

// 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库。
type Logger interface {
	Info(string)
	Close()
}

type FileLogger struct {
	filename string
	wg       sync.WaitGroup
}

func newFileLogger(filename string) *FileLogger {
	instance := &FileLogger{filename: filename}
	return instance
}

func (fl *FileLogger) Close() {
	fl.wg.Wait()
	fmt.Println("fileLogger closed")
}

func (fl *FileLogger) Info(msg string) {
	fl.wg.Add(1)
	go func() {
		defer fl.wg.Done()
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
	}()
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func newConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{}
}

type ConsoleLogger struct {
	wg sync.WaitGroup
}

func (cl *ConsoleLogger) Close() {
	cl.wg.Wait()
	fmt.Println("ConsoleLogger closed")
}

func (cl *ConsoleLogger) Info(msg string) {
	cl.wg.Add(1)
	go func() {
		defer cl.wg.Done()
		fmt.Println(msg)
	}()
}

func homework2() {
	var logger Logger
	fileLogger := newFileLogger("log.txt")
	logger = fileLogger
	defer logger.Close()
	for i := 0; i < 100; i++ {
		logger.Info(fmt.Sprintf("Hello%d", i))
	}

	consoleLogger := newConsoleLogger()
	logger = consoleLogger
	defer logger.Close()
	for i := 0; i < 100; i++ {
		logger.Info(fmt.Sprintf("Hello%d", i))
	}
}

func main() {
	homework1()
	println()
	homework2()
}
