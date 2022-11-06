package main

import (
	"fmt"
	"time"
)

func timeIt(name string) func() {
	startTime := time.Now()

	return func() {
		duration := time.Since(startTime)
		fmt.Printf("%s took %s", name, duration)
	}
}

func dot() {
	defer timeIt("dot")()

	for {
		if time.Now().Second()%11 == 0 {
			break
		}
	}
}

func main() {
	dot()
}
