package main

import (
	"fmt"
	"time"
)

func parse() {
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

func main() {
	parse()
}