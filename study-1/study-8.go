package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool
var x, y int = 1, 2

const Pi float64 = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	//runs()
	//defer days()
	//hours()
	deferstack()
}

func check() {
	var i int
	var m, h, n = true, false, "nil"
	fmt.Println("hello world")
	fmt.Println("当前运行版本：" + runtime.Version())
	fmt.Println(rand.Intn(13))
	fmt.Println(add(23456, 12345))
	a, b := swap("hello", "king")
	fmt.Println(a + b)
	fmt.Println(split(17))
	fmt.Println(c, python, java, i)
	fmt.Println(add(x, y))
	fmt.Println(m, h, n)
	i = 32
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, f, u)
	const name = "king"
	fmt.Println(name)
	fmt.Println(Pi)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(sum(0))
	fmt.Println(sum(20))
	fmt.Println(sumone(23))
	fmt.Println(sqrt(-4), sqrt(100))
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
	fmt.Println(math.Sqrt(99))
	fmt.Println(newSqrt(99))

}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum(x int) int {
	if x <= 0 {
		return 0
	}
	sum := 0
	for i := 0; i <= x; i++ {
		sum += i
	}
	return sum
}

func sumone(x int) int {
	sum := 1
	for sum <= x {
		sum += sum
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(a, b, c float64) float64 {
	if v := math.Pow(a, b); v < c {
		return v
	} else {
		return c
	}
}

func newSqrt(x float64) float64 {
	a := x
	b := 0.00001
	var c float64 = 0
	for math.Abs(a-c) >= b {
		c = a
		a = (a + x/a) / 2
	}
	return a
}

func runs() {
	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("windows")
	}
}

func days() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func hours() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")
	case t.Hour() < 20:
		fmt.Println("good evening")
	default:
		fmt.Println("go sleep")
	}
}

func deferstack() {
	fmt.Println("going...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
