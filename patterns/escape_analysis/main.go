package main

// 逃逸分析: https://www.bilibili.com/video/av70284152?p=2
// 栈比堆快, 因此能分配熬栈上就不分配到堆上
// go build -gcflags "-m -l" main.go
// go tool compile -S main.go
func returnAddr() (*int, int, *int) {
	b := new(int)  // 这里b未被外界引用, 因此分配到stack, returnAddr new(int) does not escape
	*b = 20
	c := new(int)  // c作为指针返回给外部, 因此分配到堆上, new(int) escapes to heap
	a := 10  // a的地址被返回, 因此被分配到堆上, moved to heap: a
	return &a, *b, c
}

func main() {

}