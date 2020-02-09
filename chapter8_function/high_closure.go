package main

import (
	"errors"
	"fmt"
)

/*
	高阶函数
*/
type Calculate func(x, y int) int

func cal(x, y int, f Calculate) int {
	return f(x, y)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func do(op string) (Calculate, error) {
	switch op {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	case "*":
		return mul, nil
	case "/":
		return div, nil
	default:
		return nil, errors.New(op)
	}
}

func anonymousFunc() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 20))

	ret := func(x, y int) int {
		return x / y
	}(23, -4)
	fmt.Println(ret)
}

func closureDemo1() {
	adder := func() func(int) int {
		var x int
		return func(y int) int {
			x += y
			return x
		}
	}

	f := adder()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 30
	fmt.Println(f(30)) // 60

	f = adder()
	fmt.Println(f(40)) // 40
	fmt.Println(f(20)) // 60
}

func closureDemo2() {
	adder2 := func(x int) func(int) int {
		return func(y int) int {
			x += y
			return x
		}
	}

	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(prefix string) string {
		return fmt.Sprintf("%s%s", prefix, suffix)
	}
}

func suffixFuncDemo() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

func doubleFuncDemo() {
	calc := func(x int) (f1, f2 func(int) int) {
		f1 = func(y int) int {
			x += y
			return x
		}
		f2 = func(y int) int {
			x -= y
			return x
		}
		return
	}
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

func main() {
	ret := cal(10, 11, func(x, y int) int {
		return x + y
	})
	fmt.Println(ret)
	println()

	f, err := do("*")
	if err == nil {
		fmt.Println(f(100, 20))
	}
	f, err = do("(")
	if err != nil {
		fmt.Println(err)
	}

	println()
	anonymousFunc()

	println()
	closureDemo1()

	println()
	closureDemo2()

	println()
	suffixFuncDemo()

	println()
	doubleFuncDemo()
}
