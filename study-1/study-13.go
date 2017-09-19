package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	//time1()
	//time2()
	//workpool()
	//limiter()
	//addunit()
	//mutex()
	getchan()
}

func getchan() {
	var ops int64
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	opsfinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsfinal)
}

func mutex() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops uint64 = 0
	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)
	opsfinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsfinal)
	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()
}

func time1() {
	// time1 := time.NewTimer(2 * time.Second)
	// <-time1.C
	// fmt.Println("timer 1 expired")
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("timer 2 expired")
	}()
	time.Sleep(time.Second * 2)
	stop := time2.Stop()
	if stop {
		fmt.Println("timer 2 stoped")
	}
}

func time2() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("ticker at ", t)
		}
	}()
	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("ticker stoped")
}

func workpool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 0; i < 9; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 9; i++ {
		fmt.Println(<-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println("id ", id, "processing job", i)
		time.Sleep(time.Second)
		results <- i
	}
}

func limiter() {
	// requests := make(chan int, 5)
	// for i := 1; i <= 5; i++ {
	// 	requests <- i
	// }
	// close(requests)
	// limiter := time.Tick(time.Second)
	// for req := range requests {
	// 	<-limiter
	// 	fmt.Println("request", req, time.Now())
	// }
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func addunit() {
	var ops uint64 = 0
	var ops1 int64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddInt64(&ops1, 1)
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	opsFinal1 := atomic.LoadInt64(&ops1)
	fmt.Println("ops:", opsFinal)
	fmt.Println("ops1:", opsFinal1)
}
