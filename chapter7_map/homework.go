package main

import (
	"strings"
	"fmt"
)

func magic() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}

func countWord() {
	words := "how do you do, I'm fine, thank you!"
	slice := strings.Split(
		strings.Replace(strings.Replace(words, ",", "", -1), "!", "", -1), 
		" ")
	wordCount := make(map[string]int, 100)	
	for _, v := range slice {
		wordCount[v]++
	}
	for k, v := range wordCount {
		fmt.Printf("%s=%d\n", k, v)
	}

}

func main() {
	countWord()
	println()
	magic()
}