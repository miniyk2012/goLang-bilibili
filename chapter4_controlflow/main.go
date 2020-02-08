package main

import "fmt"

func ifDemo() {
	
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

func forDemo() {
	var i int = 0
	for i < 3 {
		fmt.Printf("%d ", i)
		i++
	}
	println()

	for {
		if i >= 6 {
			break
		}
		fmt.Printf("%d ", i)
		i++
	}
	println()
}

func switchDemo1() {
	finger := 2
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	default: 
		fmt.Println("其他")
	}
}

func switchDemo2() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3,4,5: 
		fmt.Println("其他")
	default:
		fmt.Println("无效的输入！")
	}
}
// 类似if
func switchDemo3() {
	age := 20
	switch {
	case age < 30:
		fmt.Println("年轻")
	case age < 60 && age >= 30:
		fmt.Println("年壮")
	default:
		fmt.Println("年老")
	}
}

func gotoDemo2() {
	// 标签
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	breakTag:
		fmt.Println("结束for循环")
}

func breakDemo1() {
	BREAKDEMO1:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j == 2 {
					break BREAKDEMO1
				}
				fmt.Printf("%v-%v\n", i, j)
			}
		}
	fmt.Println("...")
}

func continueDemo() {
	forloop1:
		for i := 0; i < 5; i++ {
			// forloop2:
			for j := 0; j < 5; j++ {
				if i == 2 && j == 2 {
					continue forloop1
				}
				fmt.Printf("%v-%v\n", i, j)
			}
		}
}
func main() {
	ifDemo()
	forDemo()
	switchDemo1()
	switchDemo2()
	switchDemo3()
	gotoDemo2()
	breakDemo1()
	continueDemo()
}