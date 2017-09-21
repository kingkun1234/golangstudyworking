package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	//strFmt()
	//regexpFunc()
	//jsonFunc()
	timeFunc()
}

func timeFunc() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3:04PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
	sec := t.Unix()
	nanos := t.UnixNano()
	mills := nanos / 1000000
	p(sec)
	p(nanos)
	p(mills)
	p(time.Unix(sec, 0))
	p(time.Unix(0, nanos))
	then := time.Date(2017, 9, 21, 23, 24, 13, 233, time.UTC)
	p(then)
	p(t.Year())
	p(int(t.Month()))
	p(t.Day())
	p(t.Hour())
	p(t.Minute())
	p(t.Second())
	p(t.Nanosecond())
	p(t.Location())
	p(int(t.Weekday()))
	p(t.Before(then))
	p(t.After(then))
	p(t.Equal(then))
	diff := t.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(diff.String())
}

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func jsonFunc() {
	bolB, _ := json.Marshal(false)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(123)
	fmt.Println(string(intB))
	floB, _ := json.Marshal(1.333)
	fmt.Println(string(floB))
	strB, _ := json.Marshal("hello")
	fmt.Println(string(strB))
	slcD := []string{"apple", "banana", "orange"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[int]string{1: "hello", 2: "king"}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	rep1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "banana", "orange"}}
	rep1B, _ := json.Marshal(rep1D)
	fmt.Println(string(rep1B))
	rep2D := Response2{
		Page:   33,
		Fruits: []string{"hello", "king"}}
	rep2B, _ := json.Marshal(rep2D)
	fmt.Println(string(rep2B))
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println(res)
	fmt.Println(res.Page)
	fmt.Println(res.Fruits[0])
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 3, "orange": 13}
	enc.Encode(d)
	enc.Encode(d)
	fmt.Println(enc)
	fmt.Println(d)
}

func regexpFunc() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))
	a := []byte("peach punch pinch")
	fmt.Println(r.FindAll(a, -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

type data struct {
	a, b int
}

func strFmt() {
	var p = fmt.Printf
	d := data{1, 2}
	p("%v\n", d)
	p("%+v\n", d)
	p("%#v\n", d)
	p("%T\n", d)
	p("%t\n", false)
	p("%d\n", 123)
	p("%b\n", 14)
	p("%c\n", 33)
	p("%x\n", 456)
	p("%f\n", 78.9)
	p("%e\n", 123400000.0)
	p("%E\n", 123400000.0)
	p("%s\n", "\"string\"")
	p("%q\n", "\"string\"")
	p("%x\n", "hex this")
	p("%p\n", &d)
	p("|%6d|%6d|\n", 123, 45)
	p("|%6.2f|%6.2f|\n", 1.2, 3.45)
	p("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	p("|%6s|%6s|\n", "foo", "b")
	p("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
