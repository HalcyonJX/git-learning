package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var smr studentMgr

func showMenu() {
	fmt.Print(`
     1.查看所有学生
     2.增加学生
     3.修改学生信息
     4.删除学生
     5.退出系统
 `)
}
func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student),
	}
	fmt.Println("--------欢迎来到学生管理系统------------")
	for {
		showMenu()
		fmt.Print("请输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			smr.showAllStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("滚~~~")
		}
	}
}
