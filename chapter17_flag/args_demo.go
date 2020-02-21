package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%v(%T)\n", i, v, v)
	}
}