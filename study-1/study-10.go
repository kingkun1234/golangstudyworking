package main

import (
	"fmt"
	"math"
)

func main() {
	//createMap()
	//updateMap()
	//fmt.Println(strings.Fields("hello world"))
	//getfunc()
	//testAdder()
	//testfibonacci()
	//testAbs()
	//testfabs()
	testVer()
}

type myfloat float64

func testfabs() {
	f := myfloat(math.Sqrt2)
	fmt.Println(f.fabs())
}

func (f myfloat) fabs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Ver struct {
	Lat, Lon float64
}

func testVer() {
	v := &Ver{2, 3}

	fmt.Println(v.abs())
	v.scale(2)
	fmt.Println(v.abs())
}

func (v *Ver) scale(f float64) {
	v.Lat = v.Lat * f
	v.Lon = v.Lon * f
}

func testAbs() {
	v := &Ver{2, 3}
	fmt.Println(v.abs())
}

func (v *Ver) abs() float64 {
	return v.Lat*v.Lat + v.Lon*v.Lon
}

var m map[string]Ver

var m1 = map[string]Ver{
	"google":   Ver{21.2, 23.3},
	"facebook": Ver{43.2, 22.5},
}

func createMap() {
	m = make(map[string]Ver)
	m["hello"] = Ver{
		Lat: 23.4,
		Lon: 32.1,
	}
	fmt.Println(m["hello"])
	fmt.Println(m1)
}

func updateMap() {
	m2 := map[int]string{
		1: "king",
		2: "kun",
		3: "haha",
		4: "lwlw",
	}
	fmt.Println(m2)
	delete(m2, 2)
	fmt.Println(m2)
	data, ok := m2[2]
	if ok {
		fmt.Println("m2[2] is yes", data)
	} else {
		fmt.Println("no")
	}
}

func getfunc() {
	fn := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(fn(12, 32))
	fmt.Println(compute(fn))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(12, 32)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func testAdder() {
	pos, neg := adder(), adder()
	for i := 0; i <= 10; i++ {
		fmt.Println(pos(i), neg(-i*2))
	}
}

func testfibonacci() {
	fn := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(fn())
	}
}

func fibonacci() func() int {
	index := -1
	f1 := 0
	f2 := 0
	return func() int {
		index++
		v := 0
		if index == 0 {
			v = 0
		} else if index == 1 {
			v = 1
		} else {
			v = f1 + f2
		}
		f1, f2 = f2, v
		return v
	}
}
