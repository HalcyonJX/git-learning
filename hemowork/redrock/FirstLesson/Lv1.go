package main

import (
	"fmt"
)

func main() {
	//回文数
	var palindrome string
	fmt.Println("请输入一个字符串")
	fmt.Scanf("%s\n", &palindrome)
	ss := []rune(palindrome)
	i := 0
	for ; i <= len(ss)-1-i; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			fmt.Println("不是回文串")
			return
		} else {
			fmt.Print(string(ss[i]))
		}
	}

}
