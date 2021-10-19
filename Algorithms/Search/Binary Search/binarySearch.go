package main

import "fmt"

func binarySearch(array []int, value int) int {

	low := 0
	high := len(array) - 1

	for low <= high {
		middle := (low + high) / 2

		if value == array[middle] {
			return array[middle]
		} else if value < array[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	fmt.Println(value, "not found")
	return -1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(numbers, 5))
}
