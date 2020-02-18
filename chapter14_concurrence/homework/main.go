package main

import (
	"fmt"
	"math/rand"
)

const NUM = 100
const WORKER = 5

type Data struct {
	num     int64
	total   int64
	routine int
}

func (d Data) String() string {
	return fmt.Sprintf("sum(%d)=%d, worker=%d", d.num, d.total, d.routine)
}

func generate(job chan<- int64) {
	for i := 0; i < NUM; i++ {
		job <- rand.Int63n(1000)
	}
	close(job)
}

func calculate(job <-chan int64, result chan<- Data, jobNo int) {
	for num := range job {
		total := calNumSum(num)
		result <- Data{num, total, jobNo}
	}
}

func printResults(result <-chan Data) {
	for i := 0; i < NUM; i++ {
		data := <-result
		fmt.Printf("%d: %v\n", i, data)
	}
}

func calNumSum(num int64) int64 {
	var ret int64
	for num > 0 {
		ret += num % 10
		num = num / 10
	}
	return ret
}

func homework1() {
	jobChan := make(chan int64)
	resultChan := make(chan Data)
	for i := 0; i < WORKER; i++ {
		go calculate(jobChan, resultChan, i)
	}
	go generate(jobChan)

	printResults(resultChan)
}

func homework2() {

}
func
main() {
	homework1()
	println()
	homework2()
}
