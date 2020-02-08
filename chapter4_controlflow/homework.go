package main

import "fmt"


func printMulTable(n int) {
	for line :=1; line <= n; line++ {
		printLine(line)
		println()
	}
}

func printLine(line int) {
	for i := 1; i <= line; i++ {
		fmt.Printf("%d*%d=%-2d ", i, line, line * i)
	}
}


func main() {
	printMulTable(9)
}