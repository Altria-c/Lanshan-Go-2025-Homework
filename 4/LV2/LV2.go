package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type timestampWriter struct {
	logFile io.Writer
}

func newTimestampWriter(w io.Writer) *timestampWriter {
	return &timestampWriter{logFile: w}
}

func (tw *timestampWriter) Write(p []byte) (int, error) {
	p1 := string(p)
	var un = time.Now().Unix()
	uni := strconv.FormatInt(un, 10)
	var nw = time.Now().Format("2006-01-02 15:04:05")
	all := fmt.Sprintf("时间戳：%s 时间：%s  %s\n", uni, nw, p1)
	n, err := tw.logFile.Write([]byte(all))
	return n, err
}

func main() {

	file123, err := os.OpenFile("D:\\goProjects\\lansan-learning\\4\\test2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file123.Close() //关闭文件
	logWriter := newTimestampWriter(file123)
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")

	//读取文件
	file, err1 := os.ReadFile("D:\\goProjects\\lansan-learning\\4\\test2.txt")
	if err1 != nil {

		fmt.Println(err1)
		return
	}
	fmt.Println(string(file))
}
