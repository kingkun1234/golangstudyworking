package main

import "fmt"

type dog struct {
	name string
	age  int
}

type cat struct {
	name    string
	age     int
	friends []string
}

func call_cat() {
	friends := []string{"smallred", "smallblack"}
	cat := cat{name: "mimi", age: 2, friends: friends}
	newcat := cat
	//newcat.friends = append(newcat.friends, "smallblue")
	fmt.Println(cat.friends)
	fmt.Println(newcat.friends)
}

func second_cat() {
	cat1 := cat{"hello", 4, []string{"ok1", "ok2"}}
	cat2 := cat1
	cat2.friends = make([]string, len(cat1.friends))
	copy(cat2.friends, cat1.friends)
	cat2.friends = append(cat2.friends, "ok3")
	fmt.Println(cat1)
	fmt.Println(cat2)
}

func main() {
	dog := dog{name: "wangcai", age: 3}
	newDog := dog
	if dog == newDog {
		fmt.Println("dog and newdog is equal")
	} else {
		fmt.Println("dog not equal newdog")
	}
	fmt.Println(dog.name)
	fmt.Println(newDog.name)
	newDog.name = "xiaohua"
	fmt.Println(dog.name)
	fmt.Println(newDog.name)
	//call_cat()
	second_cat()
}
