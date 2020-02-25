package main

import "fmt"

func main() {
	written, err := CopyFile("./dst.txt", "./tools.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write characters:", written)
}

