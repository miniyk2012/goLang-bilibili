package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func recv(ch chan int) {
	defer wg.Done()
	x := <-ch
	fmt.Printf("接受成功 %d\n", x)
}

func demo0() {
	ch := make(chan int)
	go func() { ch <- 10 }()
	wg.Add(1)
	go recv(ch)
	wg.Wait()
}

// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，
// 两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
func demo1() {
	ch := make(chan int)
	wg.Add(1)
	go recv(ch) // 启用goroutine从通道接收值
	time.Sleep(time.Second)
	ch <- 10
	fmt.Println("发送成功")
	close(ch)
	wg.Add(1)
	go recv(ch) // 启用goroutine从通道接收值
	wg.Wait()
}

func demo2() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Printf("len=%d, cap=%d\n", len(ch), cap(ch))
	fmt.Println("发送成功")
	wg.Add(1)
	go recv(ch)
	wg.Wait()
	fmt.Printf("len=%d, cap=%d\n", len(ch), cap(ch))
}

func demo3() {
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 20
	}() // 这样不会panic
	// ch <- 10 // 这样会panic

	wg.Add(1)
	ch2 := make(chan int, 1)
	go recv(ch2)
	ch2 <- 10
	ch2 <- 20
	wg.Wait()
}

func demo4() {
	ch := make(chan int)
	close(ch)
	if _, ok := <-ch; !ok {
		fmt.Println(ok)
	}
}

func main() {
	demo0()
	println()
	demo1()
	println()
	demo2()
	println()
	demo3()
	println()
	demo4()
}
