package main

import (
	"fmt"
	tool "github.com/miniyk2012/goLang-bilibili/chapter11_package/pkg2"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Printf("%d\n", tool.Add(1, 3))
	fmt.Println("done")
}
