package main

import "fmt"

func main1() {

	mean()
}

func mean() {
	fmt.Println("Calculate mean of next numbers:")
	numbers := []int{4, 5, 6}
	fmt.Println(numbers)

	s := sum(numbers)
	fmt.Println("Mean is ", float64(s)/float64(len(numbers)))
}

func sum(numbers []int) int {
	s := 0
	for _, value := range numbers {
		s += value
	}
	return s
}
