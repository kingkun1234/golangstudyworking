package main

import (
	"fmt"
	"time"
)

func main() {
	// for1()
	// for2()
	// for3()
	// switch1()
	// switch2()
	// switch3()
	// var a [5]int
	// fmt.Println("emp", a)
	// a[2] = 100
	// fmt.Println("set", a)
	// s := make([]string, 3)
	// fmt.Println("empty ", s)
	// s[0] = "haha"
	// s[1] = "hshs"
	// s[2] = "hdhd"
	// fmt.Println(s)
	// m := make(map[string]int, 3)
	// fmt.Println(m)
	// m["a"] = 1
	// m["b"] = 2
	// fmt.Println(m)
	// delete(m, "a")
	// fmt.Println(m)
	// a, ok := m["a"]
	// fmt.Println(a, ok)
	// nums := []int{1, 2, 3, 4}
	// for i, num := range nums {
	// 	if num == 3 {
	// 		fmt.Println(i)
	// 	}
	// }
	// for i, s := range "go" {
	// 	fmt.Println(i, s)
	// }
	// fmt.Println(sum2(1))
	// fmt.Println(sum2(1, 2, 3))
	// a := []int{1, 2, 3, 4, 5}
	// fmt.Println(sum2(a...))
	fmt.Println(fact(3))
}

type 

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func sum2(nums ...int) int {
	fmt.Println(nums, "")
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func for1() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is ", i)
	}
}

func for2() {
	i := 0
	for i < 10 {
		fmt.Println("hello ", i)
		i++
	}
}

func for3() {
	i := 10
	for {
		if i%10 == 0 {
			fmt.Println("for breaking")
			break
		}
	}
}

func switch1() {
	i := 1
	switch {
	case i%2 == 1:
		fmt.Println("i == 1")
	case i%2 == 0:
		fmt.Println("i == 2")
	}
}

func switch2() {
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekday")
	default:
		fmt.Println("it's a weekend")
	}
}

func switch3() {
	day := time.Now()
	switch {
	case day.Hour() < 23:
		fmt.Println("good morning")
	case day.Hour() < 24:
		fmt.Println("good evening")
	}
}
