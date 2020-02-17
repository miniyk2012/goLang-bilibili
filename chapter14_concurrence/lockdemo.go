package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup // 互斥锁
	lock   sync.Mutex
	rwlock sync.RWMutex // 读写锁. 读写锁非常适合读多写少的场景
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func demo1() {
	x = 0
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Printf("x=%d\n", x)
}

func doLock(useRw bool, write bool) {
	switch {
	case useRw && write:
		rwlock.Lock()
	case useRw && (!write):
		rwlock.RLock()
	case !useRw:
		lock.Lock()
	}
}

func doUnlock(useRw bool, write bool) {
	switch {
	case useRw && write:
		rwlock.Unlock()
	case useRw && (!write):
		rwlock.RUnlock()
	case !useRw:
		lock.Unlock()
	}
}

func write(useRw bool) {
	doLock(useRw, true)
	x += 1
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	doUnlock(useRw, true)
	wg.Done()
}

func read(useRw bool) {
	doLock(useRw, false)
	time.Sleep(1 * time.Millisecond) // 假设写操作耗时1毫秒
	doUnlock(useRw, false)
	wg.Done()
}

func demo2(useRw bool) {
	x = 0
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(useRw)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read(useRw)
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("使用读写锁(true)or使用互斥锁(false)=%t, %v, x=%d\n", useRw, end.Sub(start), x)
}

// 读写锁
func main() {
	demo1()
	println()
	demo2(true)
	println()
	demo2(false)
}
