package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex
)

type Task struct {
	runnable func(WorkId int)
}

func main() {
	ch := make(chan Task, 10)

	for i := 0; i < 10; i++ {
		go func(WorkId int) {
			for t := range ch {
				t.runnable(WorkId)
			}
		}(i)
	}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		t1 := Task{
			runnable: func(WorkId int) {
				defer wg.Done()
				lock.Lock()
				x++
				fmt.Printf("worker%d 将x增加到了%d\n", WorkId, x)
				lock.Unlock()

			},
		}
		ch <- t1
	}
	close(ch)
	wg.Wait()
	fmt.Println("x的最终值为:", x)
}
