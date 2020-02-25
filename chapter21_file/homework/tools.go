package main

import (
	"bufio"
	"io"
	"os"
)

func send(reader *bufio.Reader, ch chan<- string) {
	defer close(ch)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			ch <- line
			return
		}
		ch <- line
	}
}

func readThenWrite(writer *bufio.Writer, ch <-chan string) (written int64) {
	defer writer.Flush()
	for line := range ch {
		num, _ := writer.WriteString(line)
		written += int64(num)
	}
	return
}

// 实现一个拷贝文件函数, 这里用管道实现
// 有更简便的方法: https://www.liwenzhou.com/posts/Go/go_file/#autoid-4-0-0
func CopyFile(dstName, srcName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	//for {
	//	line, err := reader.ReadString('\n') //注意是字符
	//	if err == io.EOF {
	//		num, _ := writer.WriteString(line)
	//		written += int64(num)
	//		writer.Flush()
	//		return written, nil
	//	}
	//	num, _ := writer.WriteString(line)
	//	written += int64(num)
	//}
	ch := make(chan string)
	go send(reader, ch)
	written = readThenWrite(writer, ch)
	return written, nil
}
