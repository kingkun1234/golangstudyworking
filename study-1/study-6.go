package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetAwardUserName(users map[string]int64) (name string) {
	size := len(users)
	awardIndex := rand.Intn(size)
	i := 0
	for userName, _ := range users {
		if i == awardIndex {
			name = userName
			return
		}
		i++
	}
	return
}

func main() {
	var users map[string]int64 = map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}
	rand.Seed(time.Now().Unix())
	awardCount := make(map[string]int)
	for i := 0; i < 1000000; i++ {
		awardName := GetAwardUserName(users)
		if count, ok := awardCount[awardName]; ok {
			awardCount[awardName] = count + 1
		} else {
			awardCount[awardName] = 0
		}
	}
	for n, c := range awardCount {
		fmt.Printf("%v,%v\n", n, c)
	}
}
