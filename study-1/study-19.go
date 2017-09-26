package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	//syncfunc()
	mapfunc()
}

func mapfunc() {
	data := NewSynchronizedMap()
	data.Put(1, 123)
	data.Put(2, 456)
	data.Put(3, 789)
	data.Each(func(k, v interface{}) {
		fmt.Println(k, "is", v)
	})
	fmt.Println(data.Get(2))
	data.Delete(1)
	data.Each(func(k, v interface{}) {
		fmt.Println(k, "is", v)
	})
}

type SynchronizedMap struct {
	rw   *sync.RWMutex
	data map[interface{}]interface{}
}

func (sm *SynchronizedMap) Put(k, v interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	sm.data[k] = v
}

func (sm *SynchronizedMap) Get(k interface{}) interface{} {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	return sm.data[k]
}

func (sm *SynchronizedMap) Delete(k interface{}) {
	sm.rw.Lock()
	defer sm.rw.Unlock()
	delete(sm.data, k)
}

func (sm *SynchronizedMap) Each(f func(interface{}, interface{})) {
	sm.rw.RLock()
	defer sm.rw.RUnlock()
	for k, v := range sm.data {
		f(k, v)
	}
}

func NewSynchronizedMap() *SynchronizedMap {
	return &SynchronizedMap{
		rw:   new(sync.RWMutex),
		data: make(map[interface{}]interface{}),
	}
}

var count int
var wg sync.WaitGroup
var rw sync.RWMutex

func syncfunc() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go reader(i)
	}
	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}

func write(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取...\n", n)
	v := count
	fmt.Printf("读goroutine %d 读取结束，读取到的值为：%d\n", n, v)
	wg.Done()
	rw.RUnlock()
}

func reader(n int) {
	rw.RLock()
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
	rw.RUnlock()
}
