package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func demo() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n > 10 {
			break
		}
	}
}

func demo2() {
	d := time.Now().Add(time.Millisecond * 40)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer func() {
		cancel()
		fmt.Println("cancel")
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}

func demo3() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	go worker(&wg, ctx)
	time.Sleep(time.Second * 5)
	wg.Wait()
	fmt.Println("done")
}

func worker(wg *sync.WaitGroup, ctx context.Context) {
	defer func() {
		fmt.Println("worker done!")
		wg.Done()
	}()
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
}

func main() {
	demo()
	println()
	demo2()
	println()
	demo3()
}
