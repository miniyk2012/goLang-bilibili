package main

import (
	"fmt"
	"time"
)


type TimeElapseCalculator struct {
	f func(time.Duration)
}

func (c TimeElapseCalculator) Run(a time.Duration) time.Duration {
	start := time.Now()
	c.f(a)
	end := time.Now()
	return end.Sub(start)
}
func timeFun(a time.Duration) {
	time.Sleep(a * time.Second)
}

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	costTime := TimeElapseCalculator{timeFun}.Run(1)
	fmt.Println(costTime)
}
