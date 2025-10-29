package main

import "fmt"

func average(sum, count int) float64 {
	return float64(sum) / float64(count)
}

func main() {
	var (
		sum   int = 0
		count int = 0
	)

	for {
		fmt.Print("请输入一个整数(输入0结束):")
		var input int
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		sum = sum + input
		count++
	}
	fmt.Println("总和为：", sum)
	fmt.Println("个数为:", count)

	Average := average(sum, count)
	fmt.Println("平均值为:", Average)

	if Average >= 60 {
		fmt.Printf("平均成绩为%v,成绩合格", Average)
	} else {
		fmt.Printf("平均成绩为%v,成绩不合格", Average)
	}
}
