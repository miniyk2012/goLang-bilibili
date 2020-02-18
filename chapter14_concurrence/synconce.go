package main

import (
	"fmt"
	"sync"
	"time"
)

const NUM = 10

var icons map[string]int
var wg sync.WaitGroup
var lock sync.Mutex
var iconOnce sync.Once
var singletonOnce sync.Once

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
	iconOnce.Do(loadIcons) // loadIcons只能调用一次
	return icons[name]
}

// sync.Once既能加锁, 也保证只执行一次
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

type singleton struct {
}

func (s singleton) String() string {
	return "I'm a singleton"
}


var instance *singleton

// 并发安全的单例模式
func GetInstance(wg *sync.WaitGroup, c chan *singleton) {
	defer wg.Done()
	singletonOnce.Do(func() {   // 这里不能再使用iconOnce, 不同的函数都要使用不同的sync.Once
		fmt.Println("create a instance")
		instance = &singleton{}
	})
	c <- instance
}

// sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，
// 而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
func singletonDemo() {
	c := make(chan *singleton)
	wg.Add(NUM)
	for i := 0; i < NUM; i++ {
		go GetInstance(&wg, c)
	}
	for i:=0; i < NUM; i++ {
		ins := <- c
		fmt.Println(ins)
	}
	wg.Wait()
}

func main() {
	concurrentLoad1()
	println()
	concurrentLoad2()
	println()
	syncOneConcurrentLoad()
	println()
	singletonDemo()
}
