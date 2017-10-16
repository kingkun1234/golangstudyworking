package main

import (
	"fmt"
	"sync"
)

func main() {
	//getSyncCount()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	f5()
	fmt.Println(f6())
}

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (result int) {
	result = 11111
	defer func() {
		fmt.Println(result)
		result += 5
	}()
	return 100
}

func f3() (result int) {
	defer func(result int) {
		result += 5
	}(result)
	return 1
}

func f4() int {
	result := 122
	defer func() {
		result += 123
	}()
	return 3
}

func f5() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f6() (result int) {
	t := 5
	defer func() {
		t += 123
	}()
	return t
}

var wg sync.WaitGroup

//var mutex sync.Mutex

type Count struct {
	Value int
	sync.Mutex
}

func getSyncCount() {
	wg.Add(1000)
	count := &Count{Value: 0}
	for i := 0; i < 1000; i++ {
		go getCount(count)
	}
	wg.Wait()
	fmt.Println(count.Value)
}

func getCount(count *Count) {
	count.Lock()
	defer count.Unlock()
	count.Value++
	wg.Done()
}

// func getSyncCount() {
// 	wg.Add(1000)
// 	mutex := &sync.Mutex{}
// 	count := &Count{0}
// 	for i := 0; i < 1000; i++ {
// 		go getCount(count, mutex)
// 	}
// 	wg.Wait()
// 	fmt.Println(count.Value)
// }

// func getCount(count *Count, mutex *sync.Mutex) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	count.Value++
// 	wg.Done()
// }

// func getSyncNum() {
// 	wg.Add(1000)
// 	count := &Count{0}
// 	for i := 0; i < 1000; i++ {
// 		go getCount(count)
// 	}
// 	wg.Wait()
// 	fmt.Println(count.Value)
// }

// func getCount(count *Count) {
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	count.Value++
// 	wg.Done()
// }
