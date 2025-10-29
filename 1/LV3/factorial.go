package main

import "fmt"

func factorial(n int) int {
	resuit := 1
	for i := n; i > 0; i-- {
		resuit = resuit * i
	}
	return resuit
}

func main() {
	a := 5
	fmt.Println(factorial(a))
}
