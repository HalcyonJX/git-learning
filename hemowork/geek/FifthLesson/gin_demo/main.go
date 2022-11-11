package main

import (
	"awesomeProject/hemowork/geek/FifthLesson/gin_demo/api"
	"awesomeProject/hemowork/geek/FifthLesson/gin_demo/dao"
)

func main() {
	dao.ReadUser()
	api.InitRouter()
}
