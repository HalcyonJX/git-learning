package main

import (
	"fmt"
	"os"
)

type hero interface {
	attack()
	skill()
	use()
}
type yasuo struct {
	name string
}
type riven struct {
	name string
}

func (y yasuo) attack() {
	fmt.Println("哈塞给")
}
func (y yasuo) skill() {
	fmt.Println("狂风绝息斩")
}
func (y yasuo) use() {
	fmt.Println("电刀+无尽")
}
func (r riven) attack() {
	fmt.Println("断剑重铸之日，骑士归来之时")
}
func (r riven) skill() {
	fmt.Println("折翼之舞")
}
func (r riven) use() {
	fmt.Println("黑切+死亡之舞")
}
func testMethod(h hero) {
	menu1()
	for {
		var choice int
		fmt.Print("请输入你要进行的操作：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			h.attack()
		case 2:
			h.skill()
		case 3:
			h.use()
		case 4:
			os.Exit(0)
		}
	}
}
func menu1() {
	fmt.Print(`       1.使用英雄进行攻击
       2.使用英雄技能
       3.购买装备
       4.退出游戏
`)
}
func menu2() {
	fmt.Print(`
       1.英雄列表
       2.选择英雄
       3.退出游戏
`)
}
func main() {
	var r = riven{name: "锐雯"}
	var y = yasuo{"亚索"}
	fmt.Println("欢迎来到召唤师峡谷，敌军还有30s到达战场，碾碎他们")
	menu2()
	for {
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("1.", r.name)
			fmt.Println("2.", y.name)
		case 2:
			fmt.Print("请输入你的选择：")
			var option1 int
			fmt.Scanln(&option1)
			switch option1 {
			case 1:
				fmt.Println("你的选择是：", r.name)
				testMethod(r)
			case 2:
				fmt.Println("你的选择是：", y.name)
				testMethod(y)
			}
		case 3:
			os.Exit(0)
		}
	}
}
