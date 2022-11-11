package main

import (
	"fmt"
	"os"
)

type movie struct {
	name      string
	actors    []string
	movieType string
	nation    string
}

func main() {
	var option int
	menu()
	m := movie{
		name: "西线无战事",
		actors: []string{
			"费力克斯", "舒赫", "亚伦",
		},
		movieType: "战争/剧情/动作",
		nation:    "德国/美国",
	}
	for {
		fmt.Print("请输入你的选择：")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("片名：", m.name)
		case 2:
			fmt.Println("主演：", m.actors)
		case 3:
			fmt.Println("电影类型：", m.movieType)
		case 4:
			fmt.Println("制片国家：", m.nation)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("不存在这个选项！！！")
		}
	}
}
func menu() {
	fmt.Print(`
      1.电影名字
      2.电影主演
      3.电影类型
      4.制片国家
      5.退出
`)
}
