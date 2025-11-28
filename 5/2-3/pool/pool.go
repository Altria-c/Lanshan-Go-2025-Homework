package pool

import "sync"

type Task func()

type Pool struct {
	taskchan chan Task
	wg       sync.WaitGroup
}

// 创建通道worker
func NewPool(chancap int, numworker int) *Pool {
	p := &Pool{
		taskchan: make(chan Task, chancap),
	}
	p.wg.Add(numworker)
	for i := 0; i < numworker; i++ {
		go p.Work()
	}
	return p
}

// work
func (p *Pool) Work() {
	defer p.wg.Done()
	for task := range p.taskchan {
		task()
	}
}

//发送任务
func (p *Pool) Submit(task Task) {
	p.taskchan <- task
}

//等待任务结束
func (p *Pool) Wait() {
	close(p.taskchan)
	p.wg.Wait()
}
