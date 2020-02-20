package main

import (
	"github.com/gin-gonic/gin"
	"strings"
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
