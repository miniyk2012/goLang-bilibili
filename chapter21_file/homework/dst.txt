package main

import "fmt"

func main() {
	written, err := CopyFile("./dst.txt", "./src.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write characters:", written)
}

