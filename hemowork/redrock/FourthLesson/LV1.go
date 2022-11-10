package main

import (
	"fmt"
	"os"
)

type skills struct {
	heroName  string
	skillName []string
}
type maneger struct {
	allSkill map[string]skills
}

var model = make([]string, 10)
var hmr maneger

func main() {
	fmt.Println("欢迎来到召唤师峡谷！！！")
	hmr = maneger{
		allSkill: make(map[string]skills),
	}
	for {
		showMenu()
		var choice int
		fmt.Print("请输入你的选择：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			hmr.showAllHero()
		case 2:
			hmr.addHero()
		case 3:
			hmr.deleteHero()
		case 4:
			editModel(model)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("不存在这个选项！！！1")
		}
	}
}

// 功能清单
func showMenu() {
	fmt.Print(`
      1.展示所有英雄并选择一名英雄
      2.增加英雄
      3.删除英雄
      4.设置技能释放模板
      5.退出游戏
`)
}

// 展示所有英雄以及后续操作
func (h maneger) showAllHero() {
	for _, v := range h.allSkill {
		fmt.Printf("英雄：%s 技能：%s\n", v.heroName, v.skillName)
	}
	fmt.Print("选择英雄：")
	var pickHero string
	fmt.Scanln(&pickHero)
	fmt.Println("您选择的英雄是：", pickHero)
	fmt.Println("此英雄技能如下：")
	sk := h.allSkill[pickHero].skillName
	for k, v := range sk {
		fmt.Printf("%d :%s\n", k+1, v)
	}
	var choice int
	fmt.Print("请输入你要选择的技能序号：")
	fmt.Scanln(&choice)
	s := sk[choice]
	fmt.Println("释放技能模板如下：")
	for k, v := range model {
		fmt.Printf("%d :%s\n", k+1, v)
	}
	var choice1 int
	fmt.Print("请输入你要选择的模板序号：")
	fmt.Scanln(&choice1)
	m := model[choice1]
	releaseSkill(s, func(skillName string) {
		fmt.Println(m, skillName)
	})
}

// 增加英雄
func (h maneger) addHero() {
	//获取玩家输入的英雄名字和技能信息
	fmt.Println("请输入英雄名字：")
	var name string
	fmt.Scanln(&name)
	var skillName = make([]string, 10)
	for i := 0; i <= len(skillName); i++ {
		var tmp string
		fmt.Print("请输入英雄的技能(输入no结束输入英雄技能)：")
		fmt.Scanln(&tmp)
		if tmp == "no" {
			break
		}
		skillName[i] = tmp
	}
	//将信息放入map
	newHero := skills{
		heroName:  name,
		skillName: skillName,
	}
	h.allSkill[name] = newHero
	fmt.Println("添加成功")
}

// 删除英雄
func (h maneger) deleteHero() {
	//输入英雄名词查询英雄
	fmt.Print("请输入你要删除的英雄：")
loop:
	var name1 string
	fmt.Scanln(&name1)
	_, ok := h.allSkill[name1]
	if !ok {
		fmt.Println("没有这个英雄，请重新输入")
		goto loop
	}
	delete(h.allSkill, name1)
	fmt.Println("删除成功")
}

// 设置技能释放模板
func editModel(model []string) {
	for i := 0; i < len(model); i++ {
		fmt.Print("请输入技能释放模板(输入no结束录入)：")
		var temp string
		fmt.Scanln(&temp)
		if temp == "no" {
			break
		}
		model[i] = temp
	}
	fmt.Println("设置成功")
}

// 释放技能
func releaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
