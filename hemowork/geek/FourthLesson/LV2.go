package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("提醒功能如下：")
	fmt.Println("1.单次日程提醒功能")
	fmt.Println("2.重复日程提醒功能")
	fmt.Println("请输入序号：")
loop:
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		singleReminder()
		break
	case 2:
		repeatReminder()
		break
	default:
		fmt.Println("不存在这个序号，请重新选择")
		goto loop

	}
}

// 单次提醒
func singleReminder() {
	var content string
	//fmt.Println("请输入提醒时间：(格式:2006-01-02 15:04:05)")
	//不会录入时间  寄  还是直接录入提醒内容吧
	fmt.Println("请输入提醒的内容")
	fmt.Scan(&content)
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-timer.C:
		fmt.Println(content)
	}
}

// 重复提醒
func repeatReminder() {
	var content string
	fmt.Println("请输入提醒内容")
	fmt.Scan(&content)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Println(content)
	}
}
