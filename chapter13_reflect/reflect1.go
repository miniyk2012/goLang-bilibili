package main

import (
	"fmt"
	"reflect"
)

func reflectType(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type:%15v\t", t)
	// v := reflect.ValueOf(a)
	// fmt.Printf("value:%10v\t", v)
	// Kind()指底层类型
	fmt.Printf("Name:%12v\tKind:%15v\n", t.Name(), t.Kind())
}

func demo1() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

type myInt int64

func demo2() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType(d)
	reflectType(e)
	var f []int
	var g [4]int
	var h map[string]int
	var i interface{} = f
	reflectType(f)
	reflectType(g)
	reflectType(h)
	reflectType(i)
	if reflect.TypeOf(f).Kind() == reflect.Slice {
		fmt.Println("f is a Slice.")
	}

}

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	switch t.Kind() {
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Slice:
		fmt.Printf("type is slice, value is %v\n", v.Interface().([]int))
	}
	// switch x.(type) {
	// case []int:
	// 	fmt.Printf("type is slice, value is %v\n", x.([]int)[0])
	// case float32:
	// 	fmt.Printf("type is float32, value is %f\n", x)
	// case int64:
	// 	fmt.Printf("type is int64, value is %d\n", x)
	// }

}

func demo3() {
	var a float32 = 3.14
	var b int64 = 100
	var c []int = []int{1, 2, 3}
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	reflectValue(c) // type is slice, value is [1 2 3]
	reflectType(c)
	// 将int类型的原始值转换为reflect.Value类型
	d := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", d) // type c :reflect.Value
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(40)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(40)
	}
}

func setValue1(x interface{}) {
	// 听奇怪的语法, 暂时不深究
	switch z := x.(type) {
	case *int64:
		fmt.Printf("type of z is %T, value=%v\n", z, *z)
		*z = 60
	}
}

func demo4() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)

	setValue1(&a)
	fmt.Println(a)
}

func demo5() {
	// *int类型空指针
	var a *int
	fmt.Println("a == nil", a == nil)
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("a IsValid:", reflect.ValueOf(a).IsValid())
	b := struct{}{}
	fmt.Println("b IsValid:", reflect.ValueOf(b).IsValid())
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	var c int
	fmt.Println("c IsValid:", reflect.ValueOf(c).IsValid())
	fmt.Println("nil  IsValid:", reflect.ValueOf(nil).IsValid())

	// map的判断
	m := make(map[int]interface{})
	fmt.Println("不存在的键3：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
	m[3] = "xx"
	fmt.Println("存在的键3：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())

}

func main() {
	demo1()
	println()
	demo2()
	println()
	demo3()
	println()
	demo4()
	println()
	demo5()
}
