package main

import (
	"fmt"
)

func main() {
	//testTeacher()
	//testStudents()
	//testCheckName()
	testMans()
}

type mans interface {
	talk()
}

type boy struct {
	name string
}

func testMans() {
	b := boy{name: "king"}
	g := girl{boy: boy{name: "dd"}, age: 23}
	speak(&b)
	speak(&g)
}

func (b *boy) talk() {
	fmt.Println(b.name)
}

func (g *girl) talk() {
	fmt.Println(g.name)
}

func speak(m mans) {
	m.talk()
}

type girl struct {
	boy
	age int
}

type check interface {
	checkName(name string) string
}

type name struct{}

func testCheckName() {
	n := name{}
	var i check
	i = &n
	fmt.Println(i.checkName("king"))
}

func (n name) checkName(name string) string {
	if name == "king" {
		return "hello king"
	}
	return "hello"
}

type school struct {
	address string
}

type home struct {
	address string
}

type students struct {
	school
	home
	name string
}

func testStudents() {
	s := students{school{"aaa"}, home{"bbb"}, "cccc"}
	fmt.Println(s)
	s.school.getAddress()
	fmt.Println(s.school.address)
	s.home.getAddress()
	fmt.Println(s.school.address)
}

func (s *school) getAddress() {
	fmt.Println(s.address)
}

func (h *home) getAddress() {
	fmt.Println(h.address)
}

type people struct{}

type teacher struct {
	people
}

func testTeacher() {
	t := teacher{}
	t.showA()
}
func (p *people) showA() {
	fmt.Println("showA")
	p.showB()
}

func (p *people) showB() {
	fmt.Println("showB")
}

func (t *teacher) showB() {
	fmt.Println("teacher showB")
}
