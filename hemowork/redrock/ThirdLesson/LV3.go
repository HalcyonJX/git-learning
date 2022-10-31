package main

// 使用两个goroutine轮流打印100以内的数
// 比如一个goroutine打印1，另一个打印2，(即一个打印奇数，一个打印偶数)，这样打印下去到100
// 注意，要保证顺序
import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- 0
			if i%2 == 1 {
				fmt.Println("线程1打印:", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			<-ch
			//偶数
			if i%2 == 0 {
				fmt.Println("线程2打印:", i)
			}
		}
	}()
	wg.Wait()
}
