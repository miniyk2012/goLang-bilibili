package main

import (
	"fmt"
	"strconv"
	"sync"
)
const NUM int = 20
var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}


func commonMapDemo() {
	wg := sync.WaitGroup{}
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func syncMapDemo() {
	m := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main() {
	//commonMapDemo() // 会发生race condition
	println()
	syncMapDemo()
}
