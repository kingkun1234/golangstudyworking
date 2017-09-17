package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type safeCount struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *safeCount) inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *safeCount) getcount(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

type IPAdd [4]byte

func main() {
	// var a IAbs
	// f := ifloat(-math.Sqrt2)
	// a = f
	// fmt.Println(a.abs())
	// k := new(king)
	// k.x = 10
	// k.y = 10
	// // k := king{1, 2}
	// a = k
	// fmt.Println(a.abs())

	// var w Write
	// w = os.Stdout
	// fmt.Fprintf(w, "hello king\n")

	// p1 := person{"king", 22}
	// p2 := person{"kun", 32}
	// fmt.Println(p1, p2)
	// addrs := map[string]IPAdd{
	// 	"china":  {1, 2, 3, 4},
	// 	"janpan": {2, 3, 4, 5},
	// }
	// for k, v := range addrs {
	// 	fmt.Println(k, v)
	// }
	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// }

	// r := strings.NewReader("hello hello king king")
	// b := make([]byte, 8)
	// for {
	// 	n, err := r.Read(b)
	// 	fmt.Printf("n==%v,err==%v,b==%v\n", n, err, b)
	// 	fmt.Printf("b[:]==%v\n", b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	// go say("hello")
	// say("king")
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// c := make(chan int)
	// go sum(a[:5], c)
	// go sum(a[5:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)
	// c := make(chan int, 10)
	// go fibonaccia(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonaccib(c, quit)

	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Println("tick")
	// 	case <-boom:
	// 		fmt.Println("boom")
	// 		return
	// 	default:
	// 		fmt.Println("hahahah")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// }
	s := safeCount{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go s.inc("key")
	}
	time.Sleep(time.Second)
	fmt.Println(s.getcount("key"))
}

func fibonaccib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("boom")
			return
		}
	}
}

func fibonaccia(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

type myErr struct {
	When time.Time
	What string
}

type person struct {
	name string
	age  int
}

func (err *myErr) Error() string {
	return fmt.Sprintf("at %v,%s", err.When, err.What)
}

func run() error {
	return &myErr{
		When: time.Now(),
		What: "it dindn't work",
	}
}

func (p person) String() string {
	return fmt.Sprintf("%v (%vyear)", p.name, p.age)
}

type Reader interface {
	Reader(b []byte) (n int, err error)
}

type Write interface {
	Write(b []byte) (n int, err error)
}

type ReaderWrite interface {
	Reader
	Write
}

type IAbs interface {
	abs() float64
}

type ifloat float64

type king struct {
	x, y float64
}

func (f ifloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (k *king) abs() float64 {
	return math.Sqrt(k.x*k.x + k.y*k.y)
}
