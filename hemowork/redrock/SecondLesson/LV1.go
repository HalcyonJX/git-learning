package main

import "fmt"

func main() {
	a := animal{
		name:  "加菲猫",
		breed: "猫",
		age:   3,
	}
	fmt.Println(a.name, a.age, a.breed)
	fmt.Println(a.action())

}

type animal struct {
	name  string
	breed string
	age   int
}

func (a animal) action() string {
	act := "抓老鼠"
	return act
}
