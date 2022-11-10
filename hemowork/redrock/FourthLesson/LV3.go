package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 存数据的map
var database = make(map[string]string)
var mibao = make(map[string]user)

// 用户信息
type user struct {
	username string
	password string
	mibao    map[string]string
}

// 功能：1.添加用户 2.查找用户 3.查找用户密码

func main() {
	readUser()
	fmt.Println("欢迎来到腾讯QQ")
	for {
		menu()
		fmt.Print("输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("开发中，敬请期待。。。。。。。。。")
		case 4:
			os.Exit(0)
		default:
			fmt.Println("无效操作")
		}
	}
}

// 菜单
func menu() {
	fmt.Print(`
     1.登录
     2.注册
     3.找回密码
     4.退出
`)
}

// 增加用户
func addUser(username, password string) {
	database[username] = password
}

// 查找用户
func selectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

// 查找正确的密码
func selectPasswordFromUsername(username string) string {
	return database[username]
}

// 注册
func register() {
	/*注册：
	传入用户名和密码
	验证用户名是否重复，若重复则重新输入。
	注册成功，保存数据到本地。*/
	var username1, password1 string
loop:
	fmt.Print("用户名：")
	fmt.Scanln(&username1)
	fmt.Print("密码(6-16位)：")
	fmt.Scanln(&password1)
	flag := selectUser(username1)
	if flag {
		fmt.Println("此用户已被注册")
		goto loop
	}
	addUser(username1, password1)
	fmt.Println("注册成功！！")
	store(database)
}

/*
登录：
传入用户名。
验证是否有该用户，没有则重新输入。
验证密码是否正确。
*/
func login() {
	var username, password string
loop:
	fmt.Print("用户名：")
	fmt.Scanln(&username)
loop1:
	fmt.Print("密码(6-16位)：")
	fmt.Scanln(&password)
	flag := selectUser(username)
	if !flag {
		fmt.Println("用户名错误")
		goto loop
	}
	selectCorrectPassword := selectPasswordFromUsername(username)
	if password != selectCorrectPassword {
		fmt.Println("密码错误")
		goto loop1
	}
	fmt.Println("登录成功！！！")
}

// 保存数据到本地
func store(m map[string]string) {
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err:", err)
	}
	file, err1 := os.OpenFile("./user.data", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err1 != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(marshal)) //将数据先写入缓存
	writer.Flush()                      //将缓存中的内容写入文件
	content, err1 := ioutil.ReadFile("./user.data")
	if err1 != nil {
		fmt.Println("read file failed, err:", err1)
		return
	}
	json.Unmarshal(content, &database)
}

//// 找回密码
//func findPassword() {
//	fmt.Print("请输入你要找回的账号用户名：")
//	var userName string
//	fmt.Scanln(&userName)
//	fmt.Println("回答下列密保问题找回密码：")
//}

// 读文件
func readUser() {
	content, err1 := ioutil.ReadFile("./user.data")
	if err1 != nil {
		fmt.Println("read file failed, err:", err1)
		return
	}
	//fmt.Println(string(content))
	json.Unmarshal(content, &database)
}
