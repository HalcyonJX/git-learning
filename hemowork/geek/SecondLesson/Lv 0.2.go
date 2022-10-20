package main

import "fmt"

func main() {
	b := Book{Name: "校花的贴身高手"}
	w := WxGG1{Name: "wx", Age: 19, Books: []Book{b}}
	w.PrintName()
	w.PrintAge()
	w.PrintBook()
	b.PrintBookName()

}

type WxGG1 struct {
	Name  string
	Age   int
	Books []Book
}

type Book struct {
	Name string
}

func (w WxGG1) PrintName() {
	fmt.Println(w.Name)
}
func (w WxGG1) PrintAge() {
	fmt.Println(w.Age)
}

func (w WxGG1) PrintBook() {
	fmt.Println(w.Books)
}

func (b Book) PrintBookName() {
	fmt.Println(b.Name)
}
