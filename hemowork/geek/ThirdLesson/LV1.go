package main

import (
	"fmt"
	"time"
)

func main() {
	gohome(play)
}
func gohome(fun func() string) {
	fmt.Print(play())
	time.Sleep(5 * time.Second)
	fmt.Print("登dua郎")

}
func play() string {
	s := "输了啦，都是你害的"
	return s
}
