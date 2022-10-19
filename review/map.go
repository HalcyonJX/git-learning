package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	m1["zjx"] = "真的帅"
	for k, v := range m1 {
		fmt.Printf("%s:%s\n", k, v)
	}
	var a string
	fmt.Scan(&a)
	fmt.Println(a)
}
