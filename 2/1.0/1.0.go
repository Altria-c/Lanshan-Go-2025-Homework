package main

import "fmt"

func fnA(arr []int) map[int]int {
	info := make(map[int]int)

	for _, v := range arr {

		info[v]++

	}
	return info
}
func main() {
	arr1 := []int{1, 2, 4, 6, 3, 6, 2, 4, 1, 4, 1, 8, 2}

	fmt.Println(fnA(arr1))
}
