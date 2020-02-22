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

func main() {
	parseTime()
	println()
	createTimeDemo()
}
