package main

import "fmt"

func main() {
	over := make(chan bool)
	//不太懂错误的原因  感觉for放里面
	//也许是 启动goroutine要时间 for放外面 i去外面找 找到9存储进over 运行了<-over 主程序就结束了
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 9 {
				over <- true
			}
		}
	}()
	<-over
	fmt.Println("over!!!")
}
