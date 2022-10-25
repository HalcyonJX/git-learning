package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan string, 1)
var wg sync.WaitGroup

func zhuang() {
	time.Sleep(1 * time.Second)
	wg.Done()
}
func miao() {
	time.Sleep(1 * time.Second)
	wg.Done()
}
func she() {
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	for {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go zhuang()
		}
		wg.Wait()
		ch <- "装弹->"
		z := <-ch
		fmt.Print(z)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go miao()
		}
		wg.Wait()
		ch <- "瞄准->"
		m := <-ch
		fmt.Print(m)
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go she()
		}
		wg.Wait()
		ch <- "射击！！！"
		s := <-ch
		fmt.Print(s, "\n")
	}
}
