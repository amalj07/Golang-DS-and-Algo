package main

import "fmt"

func mergeSort(array []int) []int {
	length := len(array)
	if length == 1 {
		return array
	}

	middle := length / 2
	left := array[:middle]
	right := array[middle:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var sortedArray []int

	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			sortedArray = append(sortedArray, left[leftIndex])
			leftIndex++
		} else {
			sortedArray = append(sortedArray, right[rightIndex])
			rightIndex++
		}
	}
	sortedArray = append(sortedArray, left[leftIndex:]...)
	sortedArray = append(sortedArray, right[rightIndex:]...)
	return sortedArray
}

func main() {
	numbers := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	fmt.Println(mergeSort(numbers))
}
