package main

import (
	"fmt"
	"time"
)

func main2() {
	startTime := time.Now()
	for {
		if time.Now().Second()%11 == 0 {
			break
		}
	}

	fmt.Println("\n", time.Since(startTime).Seconds())
}
