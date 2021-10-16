package main

import "fmt"

func insertionSort(array []int) []int {
	length := len(array)

	for i := 1; i < length; i++ {
		for i > 0 && array[i] < array[i-1] {
			array[i], array[i-1] = array[i-1], array[i]
			i--
		}
	}
	return array
}

func main() {
	numbers := []int{23, 12, 3, 56, 16, 45, 20, 32, 8, 18, 78, 62}
	fmt.Println(insertionSort(numbers))
}
