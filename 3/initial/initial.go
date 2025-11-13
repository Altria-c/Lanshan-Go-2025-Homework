package initial

import (
	"3/userinfo"
	"fmt"
)

func Menu() {
	fmt.Println("————————————————————————————————————————————————————————————————————————————————————————")
	fmt.Println("1. 唤醒旅行者    2. 更改角色    3. 祈愿    4. 展示已拥有角色    5.冒险(未开放)    6.退出")
	fmt.Println("请选择：")
	fmt.Println("————————————————————————————————————————————————————————————————————————————————————————")
}
func Getinfo1() *userinfo.Userinfo {
	var (
		name    string
		gender  string
		element string
		level   int
	)
	fmt.Println("快醒醒，快醒醒，旅行者，你的名字是：")
	fmt.Scanln(&name)
	fmt.Println("你对自己的性别有着清晰的认知，它是：")
	fmt.Scanln(&gender)
	fmt.Println("你想拥有的元素属性是：")
	fmt.Scanln(&element)
	fmt.Println("你的等级是：")
	fmt.Scanln(&level)
	fmt.Println("那么现在开始你的冒险吧！")
	return userinfo.AnUserinfo(name, gender, element, level)
}
func Getinfo2() *userinfo.Userinfo {
	var (
		name    string
		gender  string
		element string
		level   int
	)
	fmt.Println("请输入你要更改的角色姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入更改后的性别：")
	fmt.Scanln(&gender)
	fmt.Println("请输入更改后的元素：")
	fmt.Scanln(&element)
	fmt.Println("请输入更改后的等级：")
	fmt.Scanln(&level)
	return userinfo.AnUserinfo(name, gender, element, level)
}
