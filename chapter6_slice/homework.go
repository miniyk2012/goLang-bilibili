package main

import (
	"fmt"
	"sort"
)

func h1(i int) {
	var a = make([]string, i, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%v, len=%d, cap=%d\n", a, len(a), cap(a))
}

func sortArray1() {
	var a = [...]int{3, 7, 8, 9, 1}
	b := a[:]
	fmt.Printf("before sort, %v\n", a)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Printf("after sort, %v\n", a)
}

type ByValue []int

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i] < a[j] }

func sortArray2() {
	var a = [...]int{3, 7, 8, 9, 1}
	b := a[:]
	fmt.Printf("before sort, %v\n", a)
	sort.Sort(ByValue(b))
	fmt.Printf("after sort, %v\n", a)
}

func main() {
	h1(5)
	h1(0)
	println()
	sortArray1()
	println()
	sortArray2()
}
