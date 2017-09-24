package main

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	//urlfunc()
	//countHeight()
	//sha1func()
	b64func()
}

func b64func() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

func sha1func() {
	str := "hello world"
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	fmt.Println(str)
	fmt.Printf("%x\n", bs)
}

//一球从100米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第10次落地时，共经过多少米？第10次反弹多高？
func countHeight() {
	h, s := 100, 100
	for i := 1; i <= 10; i++ {
		s += h
		h /= 2
	}
	fmt.Println("经过的路程：", s)
	fmt.Println("反弹的高度:", h)
}

type Node struct {
	Value      interface{}
	next, prev *Node
}

type List struct {
	root   Node
	length int
}

func NodeNew() *List {
	l := &List{}
	l.length = 0
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *List) IsEmpty() bool {
	return l.root.next == &l.root
}

func (l *List) Length() int {
	return l.length
}

func (l *List) PushFront(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = l.root.next
		n.prev = &l.root
		l.root.next.prev = n
		l.root.next = n
		l.length++
	}
}

func (l *List) PushBack(elements ...interface{}) {
	for _, element := range elements {
		n := &Node{Value: element}
		n.next = &l.root
		n.prev = l.root.prev
		l.root.prev.next = n
		l.root.prev = n
		l.length++
	}
}

func (l *List) Find(element interface{}) int {
	index := 0
	p := l.root.next
	for p != &l.root && p.Value != element {
		p = p.next
		index++
	}
	if p != &l.root {
		return index
	}
	return -1
}

func (l *List) Remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	l.length--
}

func (l *List) Lpop() interface{} {
	if l.length == 0 {
		return nil
	}
	n := l.root.next
	l.Remove(n)
	return n.Value
}

func (l *List) NomalIndex(index int) int {
	if index > l.length-1 {
		index = l.length - 1
	}
	if index < l.length-1 {
		index = 0
	}
	index = (index + l.length - 1) % index
	return index
}

// func (l *List) Range(start, end int) []interface{} {
// 	start = l.NomalIndex(start)
// 	end = l.NomalIndex(end)
// 	res := []interface{}{}
// 	if start > end {
// 		return res
// 	}
// 	sNode := l.index(start)
// 	eNode := l.index(end)
// 	for n := sNode; n != eNode; {
// 		res = append(res, n.Value)
// 		n = n.next
// 	}
// 	res = append(res, eNode.Value)
// 	return res
// }

func urlfunc() {
	urlstr := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(urlstr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
