package main

import "fmt"

type studentMgr struct {
	allStudent map[int64]student
}

func (m studentMgr) showAllStudent() {
	for _, v := range m.allStudent {
		fmt.Printf("学生的姓名：%s 学生的学号：%d\n", v.name, v.id)
	}
}
func (m studentMgr) addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生的学号：")
	fmt.Scanln(&id)
	newStu := student{
		id:   id,
		name: name,
	}
	//将学生增加到map中
	m.allStudent[id] = newStu
	fmt.Println("添加成功")
}
func (m studentMgr) editStudent() {
	/*获取用户输入的学号查询学生
	没有就查无此人
	有就输入新的学生信息
	更新学生名字  提示成功
	*/

	fmt.Print("请输入学生的学号：")
	var a int64
	fmt.Scanln(&a)
	stuObj, ok := m.allStudent[a]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Println("你要修改的学生信息如下：")
	fmt.Printf("学号：%d 姓名：%s\n", stuObj.id, stuObj.name)
	fmt.Print("请输入新的学生姓名：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	m.allStudent[a] = stuObj
	fmt.Println("修改成功！")
}
func (m studentMgr) deleteStudent() {
	//1.输入学生的ID
	//2.没有就提示查无此人
	//3.有的话就删除 然后提示删除成功
	fmt.Print("请输入学生的学号：")
loop:
	var deleteID int64
	fmt.Scanln(&deleteID)
	_, ok := m.allStudent[deleteID]
	if !ok {
		fmt.Println("查无此人，请重新输入学号")
		goto loop
	}
	delete(m.allStudent, deleteID)
	fmt.Println("删除成功！！！！")
}
