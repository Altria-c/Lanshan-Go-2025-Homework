package work

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 搜索文件任务
type Task struct {
	FilePath string
	Keyword  string
}

type Result struct {
	FilePath    string
	numline     int
	contentline string
}

type Pool struct {
	taskchan   chan Task
	resultchan chan Result
	wg         sync.WaitGroup
}

// 创建pool并执行任务
func NewPool(chancap int, numworker int) *Pool {
	p := &Pool{
		taskchan:   make(chan Task, chancap),
		resultchan: make(chan Result, chancap),
	}
	p.wg.Add(numworker)
	for i := 0; i < numworker; i++ {

		go p.Work()
	}
	return p
}

// 处理任务
func (p *Pool) Work() {
	defer p.wg.Done()
	for task := range p.taskchan {
		content, err := os.ReadFile(task.FilePath) //读取文件内容

		if err != nil {
			continue
		}

		altercontent := string(content)
		lines := strings.Split(altercontent, "\n") //按行分割
		for num, line := range lines {
			//检查是否包含关键字
			if strings.Contains(strings.ToLower(line), strings.ToLower(task.Keyword)) {
				p.resultchan <- Result{FilePath: task.FilePath,
					numline:     num + 1,
					contentline: line,
				}
			}
		}
	}
}

// 获取结果
func (p *Pool) GetResult() <-chan Result {
	return p.resultchan

}

// 发送任务
func (p *Pool) Submit(task Task) {
	p.taskchan <- task
}

// 等待任务结束
func (p *Pool) Wait() {
	close(p.taskchan)
	p.wg.Wait()
	close(p.resultchan)
}

// 获取文件路径
func GetFile(directory string) []string {
	var files []string
	ignor := map[string]bool{
		".exe": true,
	}

	filepath.WalkDir(directory, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if !d.IsDir() {
			extent := strings.ToLower(filepath.Ext(path))
			if !ignor[extent] {
				files = append(files, path)
			}
		}
		return nil
	})
	return files
}

// 打印
func PrintRsult(r Result) {
	fmt.Println(r.FilePath, "第", r.numline, "行:", r.contentline)
}
