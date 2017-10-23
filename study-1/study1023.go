package main

import (
	"time"
)

func main() {

}

func gen(seed int, rets []chan int) {
	for {
		for _, ch := range rets {
			ch <- seed
			time.Sleep(time.Millisecond * 50)
		}
	}
}

func makeChs(len int) (rets []chan int) {
	rets = make([]chan int, len)
	for i := 0; i < len; i++ {
		rets[i] = make(chan int)
	}
	return
}
