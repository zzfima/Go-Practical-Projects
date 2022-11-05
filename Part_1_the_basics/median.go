package main

import (
	"fmt"
	"sort"
)

func main2() {
	median([]int{5, 2, 8, 4, 9, 11, 7})
	median([]int{5, 2, 8, 4, 9, 11})

	a1 := []int{5, 2, 8, 4, 9, 11}

	f1(a1)
	fmt.Println(a1)
}

func f1(a []int) {
	sort.Ints(a)
}

func median(numbers []int) {
	fmt.Println("Calculate median of array", numbers)
	numbersLength := len(numbers)
	numbersSorted := make([]int, numbersLength)
	copy(numbersSorted, numbers)
	sort.Ints(numbersSorted)

	var medianValue = 0.0
	if numbersLength%2 != 0 {
		medianValue = float64(numbersSorted[numbersLength/2])
	} else {
		i := int(numbersLength / 2)
		medianValue = float64(numbersSorted[i-1]+numbersSorted[i]) / 2
	}

	fmt.Println("Median is", medianValue)
}
