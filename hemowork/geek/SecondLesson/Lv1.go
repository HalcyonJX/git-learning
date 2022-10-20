package main

import "fmt"

func main() {
	s1 := question{
		Id:        "金牌送货员",
		content:   "《英雄联盟》S12总决赛8强出炉：RNG VS T1，RNG能进4强吗？",
		CreatedAt: "2022-10-17 10:47:31",
		UpdateAt:  "2022-10-17 10:47:31",
		answer: answer{
			Id:        "Gentle Devil",
			content:   "已经是话题度拉满的好签了",
			CreatedAt: "2022-10-20 08:44",
			UpdateAt:  "2022-10-20 08:44",
			comments: comments{
				Id:        "啊乎仲裁委",
				content:   "健康第一位，成绩可以先放放，新冠刚开始阳影响还是挺大的，安全回来就好，是吧滔博，没用的东西！",
				CreatedAt: "10-17",
				Ip:        "山东",
			},
		},
	}
	fmt.Printf("提问者：%s\n问题：%s\n 提问时间：%s\n 提问更新时间：%s\n", s1.Id, s1.content, s1.CreatedAt, s1.UpdateAt)
	fmt.Printf("回答者：%s\n内容：%s\n提问时间：%s\n 提问更新时间：%s\n", s1.answer.Id, s1.answer.content, s1.answer.CreatedAt, s1.answer.UpdateAt)
	fmt.Println("评论者：", s1.answer.comments.Id)
	fmt.Println("评论：", s1.answer.comments.content)
	fmt.Println("时间：", s1.answer.comments.CreatedAt)
	fmt.Println("ip地址：", s1.answer.comments.Ip)
}

type question struct {
	Id        string
	content   string
	CreatedAt string
	DeletedAt string
	UpdateAt  string
	answer    answer
}
type answer struct {
	Id        string
	content   string
	CreatedAt string
	DeletedAt string
	UpdateAt  string
	comments  comments
}
type comments struct {
	Id        string
	content   string
	CreatedAt string
	DeletedAt string
	UpdateAt  string
	Ip        string
}
