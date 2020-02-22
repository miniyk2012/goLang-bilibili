package main

import (
	"fmt"
	"time"
)

func parseTime() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println(err)
		return
	}
	timeString := now.Format("2006/01/02 15:04:05")
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeString, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func createTimeDemo() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeObj := time.Date(1989, 3, 15, 22, 30, 0, 0, loc)
	fmt.Println(timeObj)
}
func timeDemo() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%v-%v-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemo() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
}

func timestampDemo2() {
	timestamp := time.Now().UnixNano()
	timeObj := time.Unix(0, timestamp) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println(timeObj.Format("2006/01/02 15:04:05"))
}

func tickDemo() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	var count int
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
		count++
		if count > 10 {
			break
		}
	}
}
func main() {
	timeDemo()
	println()
	timestampDemo()
	println()
	timestampDemo2()
	println()
	parseTime()
	println()
	createTimeDemo()
	println()
	tickDemo()
}
