package main

import (
	"fmt"
)

func main() {
	slice := []int64{213, 32, 423, 5, 342, 42455, 678, 442, 523, 556, 42341, 423098, 89677}
	fmt.Println("冒泡排序：")
	BubblingSort(slice)
	fmt.Println("选择排序：")
	SelectSort(slice)
	fmt.Println("插入排序：")
	InsertSort(slice)

}

// 冒泡排序
func BubblingSort(slice []int64) []int64 {
	slice1 := make([]int64, len(slice))
	copy(slice1, slice)
	for i := 1; i <= len(slice1)-1; i++ {
		for j := 0; j < len(slice1)-i; j++ {
			if slice1[j] >= slice1[j+1] {
				slice1[j], slice1[j+1] = slice1[j+1], slice1[j]
			}
		}
	}
	fmt.Println(slice1)
	return slice1
}

// 选择排序
func SelectSort(slice []int64) []int64 {
	slice2 := make([]int64, len(slice))
	copy(slice2, slice)
	for i := 0; i < len(slice2); i++ {
		for j := i; j < len(slice2); j++ {
			if slice2[i] > slice2[j] {
				slice2[i], slice2[j] = slice2[j], slice2[i]

			}
		}
	}
	fmt.Println(slice2)
	return slice2
}

// 插入排序
func InsertSort(slice []int64) []int64 {
	slice3 := make([]int64, len(slice))
	copy(slice3, slice)
	for i := 1; i < len(slice3); i++ {
		for j := i; j > 0; j-- {
			if slice3[i] < slice3[j-1] {
				slice3[i], slice3[j-1] = slice3[j-1], slice3[i]
			}
		}

	}
	fmt.Println(slice3)
	return slice3
}
