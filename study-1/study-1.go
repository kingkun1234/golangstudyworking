package main

import "fmt"

func main() {
	fmt.Println("hello world")
	Greeting("hello", "world", "king")
	s := []string{"who", "are", "you"}
	Greeting("hello", s...)
	GetNumbers(1, 2, 3)
	numbers := []int{4, 5, 6}
	GetNumbers(numbers...)
	GetNumbers()
	Scan()
}

func Greeting(prefix string, who ...string) {
	fmt.Println(prefix)
	for _, value := range who {
		fmt.Println(value)
	}
}

func GetNumbers(number ...int) {
	fmt.Println(number)
}

const metersToYears float64 = 1.09361

func Scan() {
	var meters float64
	fmt.Print("input meters:")
	fmt.Scanln(&meters)
	years := meters * metersToYears
	fmt.Println(years)
}
