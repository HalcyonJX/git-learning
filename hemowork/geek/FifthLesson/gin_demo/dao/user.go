package dao

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 假数据库，用 map 实现
var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}

func AddUser(username, password string) {
	database[username] = password
	Store(database)
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

// 保存数据到本地
func Store(m map[string]string) {
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

// 读文件
func ReadUser() {
	content, err1 := ioutil.ReadFile("./user.data")
	if err1 != nil {
		fmt.Println("read file failed, err:", err1)
		return
	}
	//fmt.Println(string(content))
	json.Unmarshal(content, &database)
}
