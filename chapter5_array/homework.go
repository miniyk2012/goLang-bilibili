package main

import "fmt"
func sum() {
	a := [...]int{1, 3, 5, 7, 8}
	sum := 0
    for _, num := range a {
        sum += num
    }
    fmt.Println("sum:", sum)
}

func foundSum(sum int) {
	a := [...]int{1, 3, 5, 7, 8}
	for start := 0; start < len(a); start++ {
		for end := start+1; end < len(a); end++ {
			switch {
			case a[start] + a[end] == sum:
				fmt.Printf("a[%d]+a[%d]=%d\n", start, end, sum)
			}
		}
	}
}

func main() {
	sum()
	foundSum(8)
}