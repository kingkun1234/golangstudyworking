package main

import (
	"fmt"
	"os"
	"sort"
	s "strings"
)

func main() {
	//sort1()
	//sort2()
	//panic1()
	//file()
	//combination()
	stringFunc()
}

func stringFunc() {
	var p = fmt.Println
	p("contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasprefix:", s.HasPrefix("test", "te"))
	p("hassuffix", s.HasSuffix("test", "st"))
	p("index:", s.Index("test", "s"))
	p("join", s.Join([]string{"a", "b", "c"}, "--"))
	p("repat:", s.Repeat("a", 4))
	p("replace:", s.Replace("fooooo", "o", "a", -1))
	p("replace:", s.Replace("fooooo", "o", "a", 1))
	p("split:", s.Split("a-b-c-d", "-"))
	p("tolower:", s.ToLower("ASDDFFE"))
	p("toupper:", s.ToUpper("sdfdfe"))
	//P()
	p("len:", len("adf"))
	p("char:", "hello"[1])
}

func combination() {
	fruit := []string{"apple", "bnana", "orange", "peach", "kiwi"}
	fmt.Println(Index(fruit, "apple"))
	fmt.Println(Include(fruit, "apple"))
	fmt.Println(Any(fruit, func(v string) bool {
		return s.HasPrefix(v, "p")
	}))
	fmt.Println(All(fruit, func(v string) bool {
		return s.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(fruit, func(v string) bool {
		return s.Contains(v, "e")
	}))
	fmt.Println(Map(fruit, s.ToUpper))
}

func Index(vs []string, s string) int {
	if vs == nil || len(vs) == 0 {
		return -1
	}
	for i, v := range vs {
		if v == s {
			return i
		}
	}
	return -1
}

func Include(vs []string, s string) bool {
	if vs == nil || len(vs) == 0 {
		return false
	}
	for _, v := range vs {
		if v == s {
			return true
		}
	}
	return false
}

func Any(vs []string, f func(string) bool) bool {
	if vs == nil || len(vs) == 0 {
		return false
	}
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	if vs == nil || len(vs) == 0 {
		return false
	}
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	if vs == nil || len(vs) == 0 {
		return nil
	}
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	if vs == nil || len(vs) == 0 {
		return nil
	}
	vsf := make([]string, len(vs))
	for i, v := range vs {
		vsf[i] = f(v)
	}
	return vsf
}

func sort1() {
	strs := []string{"a", "c", "b", "d"}
	sort.Strings(strs)
	fmt.Println("strings", strs)
	fmt.Println(sort.StringsAreSorted(strs))
	ints := []int{7, 3, 2, 5, 6, 88}
	sort.Ints(ints)
	fmt.Println("ints:", ints)
	fmt.Println(sort.IntsAreSorted(ints))
}

type ByLenth []string

func (b ByLenth) Len() int {
	return len(b)
}

func (b ByLenth) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByLenth) Less(i, j int) bool {
	return len(b[i]) > len(b[j])
}

func sort2() {
	fruit := []string{"apple", "bnana", "orange", "peach", "kiwi"}
	sort.Sort(ByLenth(fruit))
	fmt.Println(fruit)
}

func panic1() {
	panic("this is a err")
	_, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
}

func file() {
	f := createFile("/defer.text")
	defer closeFile(f)
	writeFile(f)
}

func createFile(s string) *os.File {
	fmt.Println("create file")
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writeing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("close file")
	f.Close()
}
