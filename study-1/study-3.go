package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Age interface{}
}

func (d Dog) Speak() string {
	return "I am Dog"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "I am Cat"
}

type Tiger struct {
}

func (t Tiger) Speak() string {
	return "I am tiger"
}

type Hanlder interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Hanlder) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

func (w welcome) Do(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

type HanlderFunc func(k, v interface{})

func (h HanlderFunc) Do(k, v interface{}) {
	h(k, v)
}

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HanlderFunc(f))
}

func main() {
	// animals := []Animal{Dog{}, Cat{}, Tiger{}}
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }
	// dog := Dog{}
	// dog.Age = "123"
	// fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	// dog.Age = 12
	// fmt.Printf("%#v %T\n", dog.Age, dog.Age)
	// dog.Age = "hello dog"
	// fmt.Printf("%#v %T\n", dog.Age, dog.Age)

	persons := make(map[interface{}]interface{})
	persons["tom"] = 20
	persons["july"] = 21
	persons["jery"] = 23
	//var w welcome = "hello"
	//Each(persons, w)
	//Each(persons, HanlderFunc(w.selfInfo))
	//EachFunc(persons, w.selfInfo)
	EachFunc(persons, selfInfo)
}
