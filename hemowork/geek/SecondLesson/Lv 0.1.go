package main

import (
	"fmt"
)

type WxGG struct {
	Name string
	Age  int
}

func main() {

	//最常见的方式
	a := WxGG{
		Name: "wxgg1",
		Age:  18,
	}

	var b WxGG
	b.Name = "wxgg2"
	b.Age = 18

	WxJJ := struct {
		Name string
		Age  int
	}{
		"wxjj2",
		18,
	}

	c := NewWxGG("wxgg tql", 18)

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", WxJJ)
	fmt.Printf("%#v\n", c)

	Wxjj1 := struct {
		Name string
		Age  int
	}{
		"wxjj2",
		18,
	}

	fmt.Println(Wxjj1)
	fmt.Println(Wxjj1.Name)
	fmt.Println(Wxjj1.Age)

}

// 构造函数
func NewWxGG(name string, age int) *WxGG {
	return &WxGG{
		Name: name,
		Age:  age,
	}
}

type People struct {
	Name string
	Age  int
}

type Banked struct {
	Name   string
	member People
}
