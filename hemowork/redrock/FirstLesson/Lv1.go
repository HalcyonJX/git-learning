package main

import (
	"fmt"
	"strconv"
)

func main() {
	//回文数
	var palindrome string
	fmt.Println("请输入一个字符串")
	fmt.Scanf("%s\n", &palindrome)
	ss := make([]rune, 0, len(palindrome))
	for _, v := range palindrome {
		ss = append(ss, v)
	}
	sl := make([]string, len(ss))
	for index := range ss {
		sl[index] = strconv.QuoteRune(ss[index])
	} //将rune类型的换成字符串类型
	i := 0
	for ; i <= len(sl)-1-i; i++ {
		if sl[i] != sl[len(sl)-1-i] {
			fmt.Println("不是回文串")
			return
		} else {
			fmt.Print(sl[i])
		} //怎么去除打印出来的''啊
	}

}
