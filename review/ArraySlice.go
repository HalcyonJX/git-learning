package main

import "fmt"

func main() {
	//数组
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i*5 + 10
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]:%d\n", i, arr[i])
	}
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	var arrLazy = [...]int{5: 6, 6: 289, 7: 61, 22: 626}
	fmt.Println(arrLazy)
	//二维数组
	//切片
	s1 := arr[2:5]
	fmt.Println(s1)
	s2 := make([]int, 10)
	for i := 0; i < len(s2); i++ {
		s2[i] = 6*i + 10
	}
	/*for i := 0; i < len(s2); i++ {
		fmt.Printf("s2[%d]:%d\n", i, s2[i])
	}*/
	for i, v := range s2 {
		fmt.Printf("s2[%d]:%d\n", i, v)
	}
	s2 = append(s2, 0, 2, 3, 4, 5, 6, 5, 3)
	fmt.Println("--------------------------------------------------------")
	for i, v := range s2 {
		fmt.Printf("s2[%d]:%d\n", i, v)
	}
	s3 := make([]int, 15)
	copy(s3, s2)
	fmt.Println(s3)
}
