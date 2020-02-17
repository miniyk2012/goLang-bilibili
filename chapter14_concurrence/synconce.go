package main

import (
	"fmt"
	"sync"
	"time"
)

const NUM = 10000

var icons map[string]int
var wg sync.WaitGroup
var lock sync.Mutex
var iconOnce sync.Once

func loadIcons() {
	fmt.Println("load once")
	icons = map[string]int{
		"left":  1,
		"up":    2,
		"right": 3,
		"down":  4,
	}
}

// Icon 被多个goroutine调用时不是并发安全的
func Icon(name string) int {
	defer wg.Done()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

func concurrentLoad1() {
	icons = nil
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go Icon("left")
	}
	wg.Wait()
}

func concurrentLoad2() {
	icons = nil
	start := time.Now()
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			Icon("right")
			lock.Unlock()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("使用lock=%v\n", end.Sub(start))
}

func newIcon(name string) int {
	defer wg.Done()
	iconOnce.Do(loadIcons)
	return icons[name]
}

func syncOneConcurrentLoad() {
	icons = nil
	start := time.Now()
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go newIcon("down")
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("使用syncOnce=%v\n", end.Sub(start))
}

func main() {
	concurrentLoad1()
	println()
	concurrentLoad2()
	println()
	syncOneConcurrentLoad()
}
