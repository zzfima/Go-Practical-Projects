package main

import (
	"fmt"
	"time"
)

func main1() {
	fmt.Println("today is business day?", IsBusinessDay(time.Now()))
	fmt.Println("Today is 2.11.2022. next business day:", NextBusinessDay(time.Date(2022, 11, 2, 11, 11, 11, 11, time.Now().Location())))
	fmt.Println("Today is 3.11.2022. next business day:", NextBusinessDay(time.Date(2022, 11, 3, 11, 11, 11, 11, time.Now().Location())))
}

// IsBusinessDay check if provided day is business day
func IsBusinessDay(t time.Time) bool {
	weekDay := t.Weekday()
	return weekDay != time.Saturday && weekDay != time.Friday
}

// NextBusinessDay get next business day
func NextBusinessDay(t time.Time) time.Time {
	day := 24 * time.Hour
	for {
		t = t.Add(day)
		if IsBusinessDay(t) {
			return t
		}
	}
}
