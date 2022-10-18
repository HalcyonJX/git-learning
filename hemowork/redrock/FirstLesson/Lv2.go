package main

import "fmt"

func main() {
	//计算器（calculator）
	slice := make([]int, 0) //声明切片来装每次运算的结果
	var (
		a     int
		cmd   string
		b     int
		chose string
	)
	//多次运算
	for i := 0; ; i++ {
		fmt.Println("输入你要计算的式子，以空格作为间隔符号：")
		fmt.Scan(&a, &cmd, &b)          //录入数据
		result := calculator(a, cmd, b) //调用运算函数
		slice = append(slice, result)   //添加结果到切片
		fmt.Printf("第%d次计算结果为：%d\n", i+1, slice[i])
		fmt.Println("输入yes/no 继续计算或者退出") //是否再继续计算
		fmt.Scan(&chose)
		if chose == "yes" {
			continue
		} else if chose == "no" {
			for j, v := range slice {
				fmt.Printf("第%d次计算结果为：%d\n", j+1, v)
			}
			break
		}
	}
}
func calculator(a int, cmd string, b int) int {
	var result int
	switch cmd {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}
	return result
}
