package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	//print()
	award()
}

func print() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	chan_n := make(chan bool)
	chan_c := make(chan bool, 1)
	done := make(chan struct{})
	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_c
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_n <- true
		}
	}()

	go func() {
		chan_seq := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
		for i := 1; i < 10; i += 2 {
			<-chan_n
			fmt.Print(chan_seq[i])
			fmt.Print(chan_seq[i+1])
			chan_c <- true
		}
		done <- struct{}{}
	}()
	chan_c <- true
	<-done
}

func GetAwardUserName(users map[string]int64) (name string) {
	sizeOfUsers := len(users)
	award_index := rand.Intn(sizeOfUsers)
	var index int
	for u_name, _ := range users {
		if index == award_index {
			name = u_name
			return
		}
		index += 1
	}
	return
}

func award() {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 5,
		"c": 13,
		"d": 100,
		"e": 23,
		"f": 45,
	}
	rand.Seed(time.Now().Unix())
	award_start := make(map[string]int64)
	for i := 0; i < 1000; i++ {
		name := GetAwardUserName(users)
		if count, ok := award_start[name]; ok {
			award_start[name] = count + 1
		} else {
			award_start[name] = 1
		}
	}
	for name, count := range award_start {
		fmt.Printf("user: %s,award count:%d\n", name, count)
	}
	return
}
