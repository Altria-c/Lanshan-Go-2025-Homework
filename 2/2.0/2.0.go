package main

import "fmt"

func calc(x float64, op string, y float64) func() float64 {
	add := func() float64 {
		return x + y
	}
	sub := func() float64 {
		return x - y
	}
	multip := func() float64 {
		return x * y
	}
	divi := func() float64 {
		if y == 0 {
			fmt.Printf("无效")
			return 0
		}
		return x / y
	}

	switch op {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return multip
	case "/":
		return divi
	default:
		return nil
	}
}
func main() {
	var (
		s, c float64
		b    string
	)
	fmt.Println("请输入：")
	fmt.Scan(&s, &b, &c)
	a := calc(s, b, c)
	if a == nil {
		fmt.Println("无效")
	}
	fmt.Println(a())
}
