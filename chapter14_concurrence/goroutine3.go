package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func judgeClose1() {
	ch1 := make(chan int)
	wg.Add(11)
	go func() {
		for i :=0 ; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			if i, ok := <- ch1; !ok {
				fmt.Println("Closed")
				wg.Done()
				break
			} else {
				fmt.Println(i)
				wg.Done()
			}
			
		}
	}()
	wg.Wait()
}

func judgeClose2() {
	ch := make(chan int)
	wg.Add(10)
	go func() {
		for i :=0 ; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}

// 类似管道拼接在一起, 数据在管道中持续流动
func judgeClose3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~10的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			// 通道关闭后再取值ok=false
			if i, ok := <- ch1; ok {
				ch2 <- i * i
			} else {
				break
			}
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

func input(ch chan<- int) {
	for i:=0; i<5; i++ {
		ch <- i
	}
	close(ch)
}
func process(ch1 <-chan int, ch2 chan<- int) {
	for i := range ch1 {
		ch2 <- 2 * i
	}
	close(ch2)
}
func sink(ch2 <-chan int) {
	for i := range ch2 {
		fmt.Println(i)
	}
}

// 单向通道
func inOutDemo() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go input(ch1)
	go process(ch1, ch2)
	sink(ch2)
}

func worker(i int, jobs <- chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d start, job=%d\n", i, job)
		results <- job * job
		fmt.Printf("worker %d end, job=%d\n", i, job)
	}
}
func workerPool() {
	jobs := make(chan int)
	results := make(chan int, 100)
	var jobNum int = 10
	var workerNum int = 4
	for i :=0; i < workerNum; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < jobNum; i++ {
		jobs <- i
	}
	close(jobs)
	for i:=0; i < jobNum; i++ {
		t := <- results
		fmt.Println(t)
	}
}

func selectDemo(capacity int) {
	ch := make(chan int, capacity)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("get", x)
		case ch <- i:
			fmt.Println("put", i)
		default: 
			fmt.Println("啥都不干")
		}
	}
}

func selectDemo2(wait bool) {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 3
	}()
	if wait {
		wg.Wait()
	}
	select {
	case x := <-ch:
		fmt.Println("get", x)
	default: 
		fmt.Println("啥都不干")
	}

}
func main() {
	judgeClose1()
	println()
	judgeClose2()
	println()
	judgeClose3()
	println()
	inOutDemo()
	println()
	workerPool()
	println()
	selectDemo(1)
	println()
	selectDemo(0)
	println()
	selectDemo2(true)
	println()
	selectDemo2(false)
}