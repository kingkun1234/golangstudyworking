package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//datafunc()
	//childfunc()
	//sayfunc()
	//typefunc()
	//serverfunc()
	testfunctype()
}

func testfunctype() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 23, 34, 55}
	fmt.Println(slice)
	odd := filter(slice, isOdd)
	fmt.Println(odd)
	event := filter(slice, isEvent)
	fmt.Println(event)
}

func filter(slice []int, fn funcType) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

type funcType func(int) bool

func isOdd(i int) bool {
	if i%2 == 0 {
		return false
	}
	return true
}

func isEvent(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

type Fn func(int, int)

func (f Fn) Servers() {
	fmt.Println("this is func1")
}

func servers(int, int) {
	fmt.Println("this is func2")
}

func serverfunc() {
	f := Fn(servers)
	f(1, 2)
	f.Servers()
}

func gofunc1(f func()) {
	go f()
}

func gofunc2(f func(interface{}), i interface{}) {
	go f(i)
}

func gofunc(f interface{}, args ...interface{}) {
	if len(args) > 1 {
		go f.(func(...interface{}))(args)
	} else if len(args) == 1 {
		go f.(func(interface{}))(args[0])
	} else {
		go f.(func())()
	}
}

func f1() {
	fmt.Println("func f1 done")
}

func f2(i interface{}) {
	fmt.Println("func f2 done", i)
}

func f3(args ...interface{}) {
	fmt.Println("func f3 done", args)
}

func typefunc() {
	gofunc1(f1)
	gofunc2(f2, 100)
	gofunc(f1)
	gofunc(f2, "hello")
	gofunc(f3, "hello", "world", 123, 23.4)
	time.Sleep(time.Second * 4)
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

func (fn appHandler) ServerHttp(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// func init() {
// 	http.Handle("/", appHandler())
// }

type act interface {
	say()
}

type man struct {
}

type animal struct {
}

func sayfunc() {
	var w act
	m := man{}
	a := animal{}
	w = &m
	w.say()
	w = &a
	w.say()
}

func (m *man) say() {
	fmt.Println("man say hello")
}

func (a *animal) say() {
	fmt.Println("animal say hello")
}

type data struct {
	val int
}

func datafunc() {
	d := &data{val: 122}
	d.set(123)
	d.show()
}

func (d *data) set(n int) {
	d.val = n
}

func (d *data) show() {
	fmt.Println(d.val)
}

type parent struct {
	val int
}

type child struct {
	parent
	num int
}

func childfunc() {
	var c child
	c = child{parent{1}, 2}
	fmt.Println(c.parent.val)
	fmt.Println(c.val)
	fmt.Println(c.num)
}
