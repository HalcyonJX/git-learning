package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("请输入自然数的个数")
	var a int
	fmt.Scan(&a)
	//定义一个切片来装录入的自然数
	s := make([]int, a)
	fmt.Println("请输入自然数")
	for i := 0; i < a; i++ {
		fmt.Scan(&s[i])
	}
	//声明一个map
	m := make(map[int]int, len(s))
	//将切片的值作为map的key，用map的value来自然数出现的个数
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	//对map的key排序
	keys := make([]int, len(s), cap(s))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(sort.IntSlice(keys))
	for _, key := range keys {
		fmt.Printf("自然数：%d,出现次数：%d\n", key, m[key])
	}
}
