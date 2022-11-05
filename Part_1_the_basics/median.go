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
	var medianValue = 0.0

	if numbersLength%2 != 0 {
		medianValue = float64(numbers[numbersLength/2])
	} else {
		i := int(numbersLength / 2)
		medianValue = float64(numbers[i-1]+numbers[i]) / 2
	}

	fmt.Println("Median is", medianValue)
}
