package main

import "fmt"

//filter returns a slice with only the values that pred(val) returned true
func filter(pred func(int) bool, values []int) []int {
	returnValues := []int{}

	for _, v := range values {
		if pred(v) {
			returnValues = append(returnValues, v)
		}
	}

	return returnValues
}

func isOdd(num int) bool {
	return num%2 != 0
}

func isEven(num int) bool {
	return !isOdd(num)
}

func main5() {
	fmt.Println("Filter challenge")
	values := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Odd values:", filter(isOdd, values))
	fmt.Println("Even values:", filter(isEven, values))
}
