package main

import (
	"fmt"
	"strconv"
)

func demo1() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}

func demo2() {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
	s3 := string(i2)
	fmt.Printf("type:%T value:%#v\n", s3, s3)
	b1, err := strconv.ParseBool("True")
	if err != nil {
		fmt.Println("can't convert to bool")
	} else {
		fmt.Printf("type:%T value:%#v\n", b1, b1) //type:bool value:true
	}
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("type:%T value:%#v\n", f, f)
	i, err := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("type:%T value:%#v\n", i, i)
	u, err := strconv.ParseUint("f", 16, 64)
	fmt.Printf("type:%T value:%#v\n", u, u)
}

func demo3() {
	v := strconv.FormatInt(15, 16)
	fmt.Println(v)
	s3 := strconv.FormatInt(-16, 16)
	fmt.Println(s3)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)
}

func demo4() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}

func main() {
	demo1()
	demo2()
	println()
	demo3()
	println()
	demo4()
}
