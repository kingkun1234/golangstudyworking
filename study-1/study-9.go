package main

import (
	"fmt"
	"strings"
)

func main() {
	//point()
	//ver()
	//printslice()
	//printbord()
	//cutslice()
	//printsliceone()
	//printrange()
	oo := make([][]int, 10, 15)
	for i := range oo {
		oo[i] = make([]int, 0, 2)
	}

	fmt.Println(oo)
}

func point() {
	i, j := 21, 63
	p := &i
	fmt.Println(p)
	fmt.Println(*p)
	i = 100
	fmt.Println(*p)
	*p = 33
	fmt.Println(i)
	p = &j
	*p = *p / 7
	fmt.Println(j)
}

type Vertex struct {
	x int
	y int
}

func ver() {
	v := Vertex{23, 24}
	fmt.Println(v)
	v.x = 100
	fmt.Println(v)
	p := &v
	p.y = 123
	fmt.Println(v)
	v1 := Vertex{x: 22}
	v2 := Vertex{}
	v3 := &Vertex{12, 13}
	fmt.Println(v1, v2, v3)

}

func printslice() {
	var s []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]==%d\n", i, s[i])
	}
}

func printbord() {
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	for i := 0; i < len(game); i++ {
		fmt.Printf("%s\n", strings.Join(game[i], " "))
	}

}

func cutslice() {
	data := []string{"a", "s", "d", "f", "g"}
	fmt.Println("data==", data)
	fmt.Println("data[1:4]==", data[1:4])
	fmt.Println("data[:3]==", data[:3])
	fmt.Println(data[:])
	fmt.Println(&data)
	data1 := data[:]
	fmt.Println(&data1)
}

func printsliceone() {
	a := make([]int, 5)
	print("a", a)
	b := make([]int, 0, 5)
	print("b", b)
	c := b[:2]
	print("c", c)
	d := c[2:5]
	print("d", d)
	var e []int
	fmt.Println(e, len(e), cap(e))
	fmt.Println(e == nil)
	print("a", a)
	a = append(a, 1)
	print("a", a)
	a = append(a, 2)
	print("a", a)
	a = append(a, 3, 4, 5)
	print("a", a)
	o := []int{5, 6, 7}
	a = append(a, o...)
	print("a", a)
}

func print(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printrange() {
	data := []int{1, 2, 4, 8, 16, 32, 64}
	for i, v := range data {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
