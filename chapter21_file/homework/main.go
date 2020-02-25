package main

import (
	"fmt"
)

func homework1() {
	written, err := CopyFile("./dst.txt", "./tools.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write characters:", written)
}

func main() {
	Cat()
}

