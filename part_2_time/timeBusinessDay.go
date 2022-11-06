package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("today is business day?", isBusinessDay(time.Now()))
	fmt.Println("Today is 2.11.2022. next business day:", nextBusinessDay(time.Date(2022, 11, 2, 11, 11, 11, 11, time.Now().Location())))
	fmt.Println("Today is 3.11.2022. next business day:", nextBusinessDay(time.Date(2022, 11, 3, 11, 11, 11, 11, time.Now().Location())))
}

func isBusinessDay(t time.Time) bool {
	weekDay := t.Weekday()
	return weekDay != time.Saturday && weekDay != time.Friday
}

func nextBusinessDay(t time.Time) time.Time {
	day := 24 * time.Hour
	for {
		t = t.Add(day)
		if isBusinessDay(t) {
			return t
		}
	}
}
