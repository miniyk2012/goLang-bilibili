package main

import (
	"fmt"
	"reflect"
)

func sliceDemo() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	var e [4]bool = [4]bool{3: true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)
	fmt.Printf("type of d:%T\n", d)
	fmt.Printf("type of e:%T\n", e)
	fmt.Printf("type of nil:%T\n", nil)
}

type Callback func(interface{})

func printSliceLenCapCallback(a interface{}, callback Callback) {
	val := reflect.ValueOf(a)
	fmt.Printf("a=%v, len(a)=%d, cap(a)=%d\n", val, val.Len(), val.Cap())
	callback(a)
}
func printSliceLenCap(a interface{}) {
	printSliceLenCapCallback(a, func(interface{}) {})
}

func compareToNil(a interface{}) {
	val := reflect.ValueOf(a)
	fmt.Printf("%v == nil: %t\n", val.Interface(), val.Interface() == nil)
}

func sliceOnArray() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	b := a[1:]
	c := a[2:4]
	d := a[3:]
	printSliceLenCap(a)
	printSliceLenCap(b)
	printSliceLenCap(c)
	printSliceLenCap(d)
	a[4] = 100
	fmt.Println(d)

	e := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	g := e[1:3]
	printSliceLenCap(g)
	g = g[:cap(g)] // 可以扩充cap哟
	printSliceLenCap(g)
}

func sliceByMake() {
	a := make([]int, 2, 5)
	printSliceLenCap(a)
}

func printCmpNil() {
	var s1 []int            //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}           //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 1, 5) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Printf("%v == nil: %t\n", s1, s1 == nil)
	fmt.Printf("%v == nil: %t\n", s2, s2 == nil)
	fmt.Printf("%v == nil: %t\n", s3, s3 == nil)
	// printSliceLenCapCallback(s1, compareToNil)
	// printSliceLenCapCallback(s2, compareToNil)
	// printSliceLenCapCallback(s3, compareToNil)
}

func sliceMethod() {
	var a [8]int
	fmt.Printf("len(a)=%d, cap(a)=%d, ptr(a)=%p, ptr(a[0])=%p\n", len(a), cap(a), &a, &a[0])
	numSlice := a[:0]
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("len(numSlice)=%d, cap(numSlice)=%d, ptr(underlyArray)=%p, ptr(numSlice[0])=%p\n", len(numSlice), cap(numSlice), numSlice, &numSlice[0])
	}
	numSlice = numSlice[:cap(numSlice)]
	fmt.Println(numSlice)
	var citySlice []string
	citySlice = append(citySlice, "北京", "上海", "南京")
	newCities := []string{"成都", "重庆"}
	fmt.Printf("%v, len(numSlice)=%d, cap(numSlice)=%d\n", citySlice, len(citySlice), cap(citySlice))
	citySlice = append(citySlice, newCities...)
	fmt.Printf("%v, len(numSlice)=%d, cap(numSlice)=%d\n", citySlice, len(citySlice), cap(citySlice))
	fmt.Println()
	c := a[:]
	b := make([]int, len(c))
	copy(b, c)
	a[0] = -100
	fmt.Printf("%v, len(c)=%d, cap(c)=%d\n", c, len(c), cap(c))
	fmt.Printf("%v, len(b)=%d, cap(b)=%d\n", b, len(b), cap(b))
}

func removeFrom(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func main() {
	sliceDemo()
	println()
	sliceOnArray()
	println()
	sliceByMake()
	println()
	printCmpNil()
	println()
	sliceMethod()
	println()
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := a[:]
	b = removeFrom(b, 2)
	fmt.Println(a)
	fmt.Printf("%v, len(b)=%d, cap(b)=%d\n", b, len(b), cap(b))

}
