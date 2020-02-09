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
}
