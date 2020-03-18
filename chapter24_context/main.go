package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker1() {
	for {
		fmt.Println("worker1")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func worker2(exitChan chan int) {
	defer wg.Done()
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second * 1)
		select {
		case <-exitChan:
			return
		default:
		}
	}
}

func worker3(ctx context.Context) {
	defer wg.Done()
	go worker4(ctx)
	for {
		fmt.Println("worker3")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Printf("worker3 %v\n", ctx.Err())
			return
		default:
		}
	}
}

func worker4(ctx context.Context) {
	defer wg.Done()
	for {
		fmt.Println("worker4")
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			fmt.Printf("worker4 %v\n", ctx.Err())
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	exitChan := make(chan int)
	wg.Add(4)
	go worker1()
	go worker2(exitChan)
	go worker3(ctx)
	time.Sleep(time.Second * 3)
	//exitChan <- 1
	close(exitChan)
	cancel()
	exit = true
	wg.Wait()
	fmt.Println("over")

}
