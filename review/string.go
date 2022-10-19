package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//前缀和后缀
	var str string = "Hello World"
	fmt.Println("str 是否有前缀He？")
	str1 := strings.HasPrefix(str, "He")
	fmt.Println(str1)
	fmt.Println("str是否有后缀ld？")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ld"))
	//字符串包含关系
	s1 := "this is sb"
	s2 := "sb"
	fmt.Println("s1是否包含了s2")
	fmt.Printf("%t\n", strings.Contains(s1, s2))
	//判断子字符串在父字符串中的位置
	a := "Hello,I am jack"
	fmt.Printf("jack在%s中的位置是：%d\n", a, strings.Index(a, "jack"))
	//字符串替换
	b := "ni hao,wo shi ikun"
	fmt.Printf("%s\n", strings.Replace(b, "ikun", "黑子", 1))
	//重复字母统计
	c := "This is your fans  and your father"
	fmt.Printf("n字母的个数：%d,i字母的个数：%d\n", strings.Count(c, "n"), strings.Count(c, "i"))
	//重复字符串
	d := "zhen de niu bi"
	e := strings.Repeat(d, 2)
	fmt.Println(e)
	//修改字符串大小
	f := "How are you? What is your name"
	fmt.Printf("小写：%s\n", strings.ToLower(f))
	fmt.Printf("大写：%s\n", strings.ToUpper(f))
	//分割字符串
	str2 := "The quick brown fox jumps over the lazy dog"
	fmt.Printf("%s\n", strings.Fields(str2))
	fmt.Printf("%s\n", strings.Split(str2, "u"))
	sl2 := strings.Fields(str2)
	fmt.Printf("%s\n", strings.Join(sl2, ";"))
	//字符串与其他类型的转换
	var adc = "666"
	var ex int
	var xo string
	ex, _ = strconv.Atoi(adc)
	fmt.Printf("%d,%T\n", ex, ex)
	xo = strconv.Itoa(ex)
	fmt.Printf("%s,%T\n", xo, xo)

}
