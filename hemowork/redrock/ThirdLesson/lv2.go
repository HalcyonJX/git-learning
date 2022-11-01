package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := os.Create("E:\\GoDevlop\\basic-code\\awesomeProject\\hemowork\\redrock\\ThirdLesson\\plan.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	file, err1 := os.OpenFile("E:\\GoDevlop\\basic-code\\awesomeProject\\hemowork\\redrock\\ThirdLesson\\plan.txt", os.O_WRONLY, 0666)
	if err1 != nil {
		fmt.Println("open file failed, err:", err1)
		return
	}
	str := "I’m not afraid of difficulties and insist on learning programming"
	file.Write([]byte(str))
	file.Close()

	file, err2 := os.Open("E:\\GoDevlop\\basic-code\\awesomeProject\\hemowork\\redrock\\ThirdLesson\\plan.txt")
	if err2 != nil {
		fmt.Println("open file failed!, err:", err2)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	n, err2 := file.Read(tmp)
	if err2 == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err2 != nil {
		fmt.Println("read file failed, err:", err2)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}
