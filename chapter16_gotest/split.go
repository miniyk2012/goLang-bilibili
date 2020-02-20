package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miniyk2012/goLang-bilibili/chapter16_gotest/handles"
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
	r.GET("/ping", handles.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080
}
