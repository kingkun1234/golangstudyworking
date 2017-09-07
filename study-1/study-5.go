package main

import "fmt"

func main() {
	value := new(int)
	fmt.Println(*value)
	modify(value)
	fmt.Println(*value)
}

func modify(value *int) {
	value = nil
	fmt.Println(&value)
}
