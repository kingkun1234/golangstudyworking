package main

import (
	"container/list"
	"fmt"
)

type student struct {
	age  int
	name string
}

func main() {
	//defer_call()
	// var score float32 = doubleScore(120)
	// fmt.Println(score)
	// studs := pase_student()
	// for k, v := range studs {
	// 	fmt.Println("key=%s,value=%v \n", k, v)
	// }
	//list_play3()
	var str string
	if str == "" || len(str) == 0 {
		fmt.Println("12")
	} else {

	}
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stud := []student{
		{age: 12, name: "king"},
		{age: 13, name: "king1"},
		{age: 14, name: "king2"},
	}
	for k, _ := range stud {
		stu := stud[k]
		m[stu.name] = &stu
	}
	return m
}

func list_play3() {
	l := list.New()
	for i := 0; i < 4; i++ {
		l.PushBack(student{age: i, name: fmt.Sprintf("hah%d", i)})
	}
	for e := l.Front(); e != nil; e = e.Next() {
		data, ok := e.Value.(student)
		if ok {
			fmt.Println(data.age)
		}
	}
}

func list_play2() {
	l := list.New()
	l.PushBack(1)
	l.PushBack("hello")
	l.PushBack([2]string{"ha", "he"})
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func list_play() {
	l := list.New()
	for i := 0; i < 6; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
	l.InsertAfter(12, l.Front())
	l.InsertBefore(11, l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l.MoveToFront(l.Back())
	l.MoveToBack(l.Front())
	l2 := list.New()
	l2.PushFront(22)
	l2.PushBackList(l)
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l2.Init()
	fmt.Println(l2.Len())
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	l.Remove(l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || source >= 100 {
			score = source
		}
	}()
	score = source * 2
	return
}
