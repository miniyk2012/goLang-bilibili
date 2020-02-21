package main

import (
	"fmt"
	"flag"
	"os"
	"time"
)

func main() {
	var delay time.Duration
	name := flag.String("name", "杨恺", "姓名")
	age := flag.Int("age", 19, "年龄")
	married := flag.Bool("married", true, "结婚否")
	flag.DurationVar(&delay ,"d", time.Hour * 10 + time.Second*58 + time.Millisecond * 40, "时间间隔")
	//解析命令行参数
	flag.Parse()
	fmt.Println(os.Args)
	fmt.Println(*name, *age, *married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
