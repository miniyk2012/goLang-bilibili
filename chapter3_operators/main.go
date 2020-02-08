package main

import "fmt"

func main() {
	var a int = 5
	if 5 > 6 || 5 < 10 {
		a += 1
		fmt.Println(a)
		fmt.Printf("%d %% 4 = %d\n", a, a%4)
	}
}
