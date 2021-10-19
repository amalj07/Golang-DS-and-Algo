package main

import "fmt"

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	left := 0
	right := len(array) - 1
	pivot := len(array) / 2

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}

	}
	array[left], array[right] = array[right], array[left]
	quickSort(array[:left])
	quickSort(array[left+1:])
	return array
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 28, 4, 0}
	fmt.Println(quickSort(numbers))
}
