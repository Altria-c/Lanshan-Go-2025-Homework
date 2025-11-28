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

func add() {
	lock.Lock()
	x = x + 1
	lock.Unlock()
	wg.Done()
}

func main() {

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)

}
