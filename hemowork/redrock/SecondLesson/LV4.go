package main

import (
	"fmt"
	"sort"
)

// 定义序列类型
type testSort []int

func main() {
	t := testSort{10, 20, 23, 4231, 321, 423, 2, 34, 564, 6423, 563, 423}
	//正序排列
	sort.Sort(t)
	fmt.Println(t)
	//倒序
	sort.Sort(sort.Reverse(t))
	fmt.Println(t)

}

// 元素个数
func (t testSort) Len() int {
	return len(t)
}

// 比较
func (t testSort) Less(i, j int) bool {
	return t[i] < t[j]
}

// 交换方式
func (t testSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
