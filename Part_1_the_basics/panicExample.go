package main

import "fmt"

func main() {
	numbers := []int{1}
	fmt.Println(safeSecondToLast(numbers))
}

func safeSecondToLast(arr []int) (i int, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return secondToLast(arr), nil
}

func secondToLast(arr []int) int {
	return arr[len(arr)-2]
}
