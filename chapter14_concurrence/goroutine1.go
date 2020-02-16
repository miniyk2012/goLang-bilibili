package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func hello1() {
	fmt.Println("Hello Goroutine!")
}

func demo1() {
	go hello1()
	time.Sleep(time.Second)
	fmt.Println("main goroutine done!")
}

func hello2(wg *sync.WaitGroup) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!")
}

func demo2() {
	var wg sync.WaitGroup
	wg.Add(1) // 启动一个goroutine就登记+1
	go hello2(&wg)
	wg.Wait() // 等待所有登记的goroutine都结束
	fmt.Println("main goroutine done!")
}

func demo3() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("main goroutine done!")
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func maxprocDemo() {
	runtime.GOMAXPROCS(6)
	go a()
	go b()
	time.Sleep(time.Second)
}

func main() {
	demo1()
	println()
	demo2()
	println()
	demo3()
	println()
	maxprocDemo()
}
