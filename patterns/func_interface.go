/**
接口型函数: https://blog.csdn.net/flysnow_org/article/details/103521045
**/

package main

import "fmt"

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type FuncHandler func(k, v interface{})

func (f FuncHandler) Do(k, v interface{}) {
	f(k, v)
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, FuncHandler(f))
}

func selfInfo(k interface{}, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)
}

