package main

import "fmt"

func main() {
	d := Dog{"uu", 3}
	c := Cat{"mm", 2}
	fmt.Println(d.name)
	fmt.Println(c.name)
	testInterface(d)
	testInterface(c)

}

type action interface {
	sound()
	eat()
}
type Dog struct {
	name string
	age  int
}
type Cat struct {
	name string
	age  int
}

func (d Dog) sound() {
	fmt.Println(d.name, "叫声：汪汪")
}
func (d Dog) eat() {
	fmt.Println(d.name, "的食物：肉")
}
func (c Cat) sound() {
	fmt.Println(c.name, "叫声：喵喵")
}
func (c Cat) eat() {
	fmt.Println(c.name, "的食物：鱼和老鼠")
}
func testInterface(action2 action) {
	action2.sound()
	action2.eat()
}
