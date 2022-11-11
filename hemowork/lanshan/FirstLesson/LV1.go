package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, l, r int
	fmt.Println("输入数组长度和起始范围，以空格作为间隔符号：")
	fmt.Scan(&n, &l, &r)
	fmt.Println("依次输入数组的元素，以逗号作间隔：")
	var s string
	fmt.Scan(&s)
	sl := strings.Split(s, ",")
	slice := make([]int, n)
	//将[]string转为[]int
	for k, v := range sl {
		slice[k], _ = strconv.Atoi(v)
	}
	new1 := BubblingSort(slice)
	for _, v := range new1[l:r] {
		fmt.Print(v, " ")
	}
}
func BubblingSort(slice []int) []int {
	slice1 := make([]int, len(slice))
	copy(slice1, slice)
	for i := 1; i <= len(slice1)-1; i++ {
		for j := 0; j < len(slice1)-i; j++ {
			if slice1[j] >= slice1[j+1] {
				slice1[j], slice1[j+1] = slice1[j+1], slice1[j]
			}
		}
	}
	return slice1
}
