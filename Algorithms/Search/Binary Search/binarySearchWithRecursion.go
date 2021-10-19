package main

import "fmt"

func binarySearch(array []int, value, low, high int) int {

	if low > high {
		return -1
	}

	middle := (low + high) / 2

	if value == array[middle] {
		return array[middle]
	} else if value < array[middle] {
		return binarySearch(array, value, low, middle)
	} else {
		return binarySearch(array, value, middle+1, high)
	}

}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(numbers, 6, 0, len(numbers)-1))
}
