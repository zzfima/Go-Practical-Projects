package main

import (
	"fmt"
	"sort"
)

func main() {
	median([]int{5, 2, 8, 4, 9, 11, 7})
	median([]int{5, 2, 8, 4, 9, 11})
}

func median(numbers []int) {
	fmt.Println("Calculate median of array", numbers)
	sort.Ints(numbers)

	numbersLength := len(numbers)
	var medianIndex = 0
	if numbersLength%2 == 0 {
		medianIndex = numbersLength/2 - 1
	} else {
		medianIndex = numbersLength / 2
	}

	fmt.Println("Median is", numbers[medianIndex])
}
