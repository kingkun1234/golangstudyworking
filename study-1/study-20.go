package main

import (
	"fmt"
	"time"
)

func main() {
	createAt := monthcountsince(time.Date(2017, 3, 27, 0, 0, 0, 0, time.UTC))
	fmt.Println(createAt)
}

func monthcountsince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) {
		createdAtTime = createdAtTime.Add(time.Hour * 24)
		nextmonth := createdAtTime.Month()
		if month != nextmonth {
			months++
		}
		month = nextmonth
	}
	return months
}
