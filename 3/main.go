package main

import (
	"fmt"
	"lansan-learning/3/initial"
	"lansan-learning/3/userinfo"
)

func main() {
	fmt.Println("——————————————————")
	fmt.Println("  原神 启动！！！")
	info := userinfo.AnUser()
	var mgr = info

	for {
		initial.Menu()
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1: //创建角色
			usr := initial.Getinfo1()

			mgr.Adduser(usr)

		case 2: //更改角色
			usr := initial.Getinfo2()
			mgr.ChangeUser(usr)

		case 3: //抽取角色
			ex := userinfo.Extrauser()
			mgr.Adduser(ex)
		case 4: //展示角色
			mgr.Showeuser()
		case 5:
			fmt.Println("前面的区域以后再来探索吧！")
		case 6: //退出
			fmt.Println("真的要退出吗  Σ( ´ﾟДﾟ`)")
			fmt.Println("是           否 ")
			fmt.Println("5!-4!=       1+1= ")
			fmt.Println("请输入答案退出：")
			for {
				var a int
				fmt.Scanln(&a)

				if a == 96 {
					fmt.Println("好吧，算你厉害(っ╥╯﹏╰╥c)")
					return
				} else if a == 2 {
					break
				} else {
					fmt.Println("不对~嘻嘻~(´・ω・`)")
					fmt.Println("请输入答案：")
					continue
				}
			}

		default:
			fmt.Println("请输入有效数字")
		}
	}
}
