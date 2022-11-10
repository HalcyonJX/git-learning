package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go tickDemo()
	go tickDemo1()
	go tickDemo2()
	wg.Wait()
}
func tickDemo() {
	defer wg.Done()
	ticker := time.Tick(30 * time.Second)
	for i := range ticker {
		fmt.Println(i)
		fmt.Println("芜湖！起飞！")
	}
}
func tickDemo1() {
	defer wg.Done()
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
	t := time.Tick(next.Sub(now))
	for i := range t {
		fmt.Println(i)
		fmt.Println("我还能再战4小时！")
	}
}
func tickDemo2() {
	defer wg.Done()
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 6, 0, 0, 0, next.Location())
	t := time.Tick(next.Sub(now))
	for i := range t {
		fmt.Println(i)
		fmt.Println("我要去图书馆开卷！")
	}
}
