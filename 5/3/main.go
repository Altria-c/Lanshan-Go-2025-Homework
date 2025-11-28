package main

import (
	"fmt"
	"lansan-learning/5/3/work"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("格式: ./catch [目录] [关键词]")
		os.Exit(1)
	}
	directory := os.Args[1]
	keyword := os.Args[2]

	files := work.GetFile(directory)
	fmt.Printf("当前一共发现%d个文件\n", len(files))

	pool := work.NewPool(100, 5)
	//发送任务
	for _, file := range files {
		pool.Submit(work.Task{
			FilePath: file,
			Keyword:  keyword,
		})
	}
	//获取结果并打印
	var rwg sync.WaitGroup
	var count int
	rwg.Add(1)
	go func() {
		defer rwg.Done()
		for result := range pool.GetResult() {
			work.PrintRsult(result)
			count++
		}
	}()
	pool.Wait()
	rwg.Wait()
	if count == 0 {
		fmt.Println("无匹配结果")
	} else {
		fmt.Printf("共发现%d处匹配\n", count)
	}
	fmt.Println("搜索完成")

}
