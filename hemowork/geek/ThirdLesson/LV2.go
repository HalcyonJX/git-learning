package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 3)

func zhuang() {
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
	}
	ch <- "装弹->"
}
func miao() {
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
	}
	ch <- "瞄准->"
}
func she() {
	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
	}
	ch <- "射击！"
}

func main() {
	for {
		go zhuang()
		go miao()
		go she()
		z := <-ch
		m := <-ch
		s := <-ch
		fmt.Print(z, m, s)
		fmt.Print("\n")
	}
	time.Sleep(2 * time.Minute)

}
