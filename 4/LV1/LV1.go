package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func WriteWithoutBuf() {
	file123, err := os.OpenFile("D:\\goProjects\\lansan-learning\\4\\test1.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file123.Close()
	start := time.Now()
	for i := 0; i < 10000; i++ {
		file123.WriteString("hello golang hello world")
	}
	end := time.Since(start)
	fmt.Println(end)
}

func WriteWithBuf() {
	file123, err := os.OpenFile("D:\\goProjects\\lansan-learning\\4\\test1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file123.Close()
	writer := bufio.NewWriter(file123)
	start := time.Now()
	for i := 0; i < 10000; i++ {
		writer.WriteString("hello golang hello world")
	}
	writer.Flush()
	end := time.Since(start)
	fmt.Println(end)
}

func main() {
	WriteWithoutBuf() //50.0614ms
	WriteWithBuf()    //0.546ms
}
