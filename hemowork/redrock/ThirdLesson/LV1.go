package main

import (
	"fmt"
	"sync"
)

var x int64
var ch1 = make(chan int64, 1)
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		x = x + 1
	}
	ch1 <- x
}
func main() {
	wg.Add(2)
	go add()
	select {
	case <-ch1:
		go add()
	}
	wg.Wait()
	fmt.Println(x)

}
