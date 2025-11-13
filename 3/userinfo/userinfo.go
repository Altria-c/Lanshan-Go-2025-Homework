package userinfo

import (
	"fmt"
	"math/rand"
)

type Userinfo struct {
	Name    string
	Gender  string
	Element string
	Level   int
}

func AnUserinfo(name, gender, element string, Level int) *Userinfo {
	return &Userinfo{
		Name:    name,
		Gender:  gender,
		Element: element,
		Level:   Level,
	}
}

type User struct {
	AllUser []*Userinfo
}

func AnUser() *User {
	return &User{
		AllUser: make([]*Userinfo, 0, 1000),
	}
}

type Usermanager interface {
	Adduser(ur *Userinfo)
	ChangeUser(ur *Userinfo)
	Extrauser() *Userinfo
	Showeuser()
}

// 创建角色
func (u *User) Adduser(ur *Userinfo) {
	u.AllUser = append(u.AllUser, ur)
	fmt.Println("哔——数据已传入天空岛")
}

// 更改角色
func (u *User) ChangeUser(ur *Userinfo) {
	for k, v := range u.AllUser {
		if ur.Name == v.Name {
			u.AllUser[k] = ur
			fmt.Println("修改成功")
			return
		}
	}
	fmt.Println("没有这个角色")
}

// 抽取角色
var Role = []*Userinfo{
	AnUserinfo("胡桃", "女", "火", 93),
	AnUserinfo("纳西妲", "女", "草", 93),
	AnUserinfo("神里绫华", "女", "冰", 90),
	AnUserinfo("绫地宁宁", "女", "水", 100),
	AnUserinfo("阿尔托莉雅", "女", "风", 97),
	AnUserinfo("篝之雾枝", "女", "冰", 94),
	AnUserinfo("雪之下雪乃", "女", "冰", 92),
	AnUserinfo("千早爱音", "女", "火", 99),
	AnUserinfo("钟离", "男", "岩", 90),
	AnUserinfo("刻晴", "女", "雷", 90),
	AnUserinfo("灯露椎", "女", "雷", 92),
	AnUserinfo("初音未来", "女", "雷", 91),
	AnUserinfo("忍野忍", "女", "冰", 90),
	AnUserinfo("桃喰莉莉香", "女", "风", 90),
}

func Extrauser() *Userinfo {
	var num = rand.Intn(11)
	ur := Role[num]
	fmt.Println("抽取成功,纠缠之缘-1")
	fmt.Println("恭喜你抽到：", ur.Name)
	return ur
}

// 展示角色
func (u *User) Showeuser() {
	if len(u.AllUser) == 0 {
		fmt.Println("目前没有角色")
	}
	for _, v := range u.AllUser {
		fmt.Printf("姓名：%s,性别：%s,元素：%s,等级；%d\n", v.Name, v.Gender, v.Element, v.Level)
	}
}
