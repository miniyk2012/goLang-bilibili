package main

import (
	"strings"
	"github.com/miniyk2012/goLang-bilibili/handles"
	"github.com/gin-gonic/gin"
)

func Split(s, sep string) (result []string) {
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
func main() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
