package main

import (
	"fmt"
	"lansan-learning/5/2-3/pool"
	"sync"
)

var x int

func main() {
	p := pool.NewPool(10, 10)

	var lock sync.Mutex
	for i := 0; i < 50; i++ {
		task := func() {
			for i := 0; i < 100; i++ {
				lock.Lock()
				x++
				lock.Unlock()
			}
		}
		p.Submit(task)

	}
	p.Wait()
	fmt.Println(x)
}
