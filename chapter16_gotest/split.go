package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miniyk2012/goLang-bilibili/chapter16_gotest/handles"
)

func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	seqLength := len(sep)
	i := strings.Index(s, sep)
	for i != -1 {
		result = append(result, s[:i])
		s = s[i+seqLength:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func Fib(n int) int {
	if n < 2  {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FastFib(n int) int {
	a, b := 0, 1
	for i:=0; i<n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	r := gin.Default()
	r.GET("/ping", handles.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
