package main

import "fmt"

func main() {
	s1 := []int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	fmt.Println("初始顺序：", s1)
	fmt.Println("冒泡排序：")
	BubblingSort(s1)
	fmt.Println("选择排序：")
	SelectSort(s1)
	fmt.Println("插入排序：")
	InsertSort(s1)
}
func BubblingSort(s []int) []int {
	//fmt.Println(s)
	for i := 1; i <= len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] >= s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println(s)
	return s
}
func SelectSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s)-1; j++ {
			if s[i] >= s[j+1] {
				s[i], s[j+1] = s[j+1], s[i]
			}
		}
	}
	fmt.Println(s)
	return s
}
func InsertSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	fmt.Println(s)
	return s

}
