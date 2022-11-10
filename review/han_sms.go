package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
}

var (
	allStudent = make(map[int64]*student, 50)
)

// 函数版学生管理系统
func main() {
	fmt.Println("欢迎来到学生管理系统！！！！")
	for {
		fmt.Print(`
           1.查看所有学生
           2.新增学生
           3.删除学生
           4.退出 
`)
		var choice int
		fmt.Print("请输入你的选择：")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚~~~")
		}
	}
}
func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学生姓名：%s 学生学号：%d\n", v.name, k)
	}
}
func addStudent() {
	//创建一个学生
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名")
	fmt.Scanln(&name)
	newstu := newStudent(id, name)
	allStudent[id] = newstu
	//追加到map里面
}
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func deleteStudent() {
	var deleteID int64
	fmt.Print("请输入需要删除学生的id:")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}
