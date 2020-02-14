package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func mapDemo1() {
	scoreMap := make(map[string]int, 3)
	scoreMap["小米"] = 100
	scoreMap["华为"] = -10
	scoreMap["苹果"] = 30
	scoreMap["抖音"] = 60
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Println(scoreMap["华为"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"name":     "杨恺",
		"password": "123456",
	}
	fmt.Println(userInfo)

	if v, ok := userInfo["password"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("password not exist")
	}
}

func mapDemo2() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Printf("score[%s]=%d\n", k, v)
	}

	delete(scoreMap, "张三")
	delete(scoreMap, "张si")
	fmt.Println(scoreMap)
}

func mapSort() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 10; i > 0; i-- {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	keys := make([]string, 0, 10)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	fmt.Println(scoreMap)
	fmt.Println(keys)
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%d\t", scoreMap[k])
	}
	println()
}

func sortValue() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 10; i > 0; i-- {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	// 构建一个反向的map, 把keys做排序, 取到value
	reversedMap := make(map[int]string, 200)
	for k, v := range scoreMap {
		reversedMap[v] = k
	}
	keys := make([]int, 0, 10)
	for key := range reversedMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Printf("%s\t", reversedMap[v])
	}
	println()
	for _, v := range keys {
		fmt.Printf("%d\t", v)
	}
	println()

	// 通过分数做排序, 用slice实现
	type Person struct {
		name  string
		score int
	}
	var people []Person
	for k, v := range scoreMap {
		people = append(people, Person{k, v})
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].score < people[j].score
	})
	fmt.Println(people)
}

// 元素为map类型的切片
func mapSlice() {
	s := make([]map[string]int, 3, 10)
	fmt.Println(s)
	m := make(map[string]int)
	s[0] = m
	s[1] = m
	fmt.Println(s)
	m["杨恺"] = 30
	fmt.Println(s)
	delete(m, "杨恺")
	fmt.Println(s)
}

// 元素为slice类型的Map
func sliceMap() {
	m := make(map[string][]string, 3)
	fmt.Println(m)
	s := make([]string, 0, 3)
	s = append(s, "a", "b", "c")
	m["杨恺"] = s
	fmt.Println(m)
	// slice复制了, 但是底层的数组是同一个
	m["王洋"] = s
	fmt.Println(m)
	s = append(s[:1], s[2:]...)
	s = append(s, "d")
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
	fmt.Println(m)
	fmt.Println(len(m["杨恺"]), cap(m["杨恺"]))

	fmt.Printf("%p\n", m["杨恺"])
	fmt.Printf("%p\n", m["王洋"])
	fmt.Printf("%p\n", s)
}

// golang 切片和struct的赋值为值拷贝，map为引用拷贝。
func assignDemo() {
	l := make([]int64, 0)
	l = append(l, 10)
	l1 := l
	l1 = append(l1, 20)
	fmt.Println(l, l1) //[10] [10 20]
	l1[0] = 50
	fmt.Println(l, l1) //[10] [50 20]

	println()
	l2 := make([]int64, 0, 5)
	l2 = append(l2, 10)
	l3 := l2
	l3 = append(l3, 20)
	fmt.Println(l2, l3) //[10] [10 20]
	l3[0] = 50
	fmt.Println(l2, l3) //[50] [50 20]
}

func assignDemo2(alloc int) {
	m := make(map[string]interface{}, 5)
	l := make([]int64, 0, alloc)
	m["hello"] = l
	l = append(l, 1)
	fmt.Println(m["hello"]) //[]
	capacity := cap(m["hello"].([]int64))
	fmt.Println(m["hello"].([]int64)[:capacity]) // 这个不好说啦
}

func assignDemo3() {
	m := make(map[string]interface{}, 0)
	l := map[string]interface{}{
		"hello": "hi",
	}
	m["hello"] = l
	l["hi"] = 1
	fmt.Println(m["hello"]) //map[hello:hi hi:1]
}

type A struct {
	B int
}

func assignDemo4() {
	m := make(map[string]interface{}, 0)
	a := &A{1}
	m["hello"] = a
	a.B = 4
	fmt.Println(*(m["hello"].(*A))) //{4}
	m["hello"].(*A).B = 6
	fmt.Println(*(m["hello"].(*A))) //{6}
	fmt.Println(*a)                 //{6}

	println()
	m1 := make(map[string]interface{}, 0)
	l := make([]int64, 0)
	m1["hello"] = &l
	l = append(l, 10)
	fmt.Println(m1["hello"])
}

func sliceModify(l []int) {
	l[0] = 100
}

func demo5() {
	l := []int{1, 2, 3}
	sliceModify(l)
	fmt.Println(l)
}

func pointerToSlice() {
	slice := []string{"a", "a"}

	func(slice *[]string) {
		(*slice)[0] = "b"
		(*slice)[1] = "b"
		*slice = append(*slice, "a")
		fmt.Println(*slice)
	}(&slice)
	fmt.Println(slice)
}

func main() {
	mapDemo1()
	println()
	mapDemo2()
	println()
	mapSort()
	println()
	sortValue()
	println()
	mapSlice()
	println()
	sliceMap()
	println()
	assignDemo()
	println()
	assignDemo2(0)
	println()
	assignDemo2(3)
	println()
	assignDemo3()
	println()
	assignDemo4()
	println()
	demo5()
	println()
	pointerToSlice()
}
