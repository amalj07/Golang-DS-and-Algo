package main

import "fmt"

func selectionSort(array []int) []int {
	length := len(array)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	return array
}

func main() {
	numbers := []int{23, 12, 3, 56, 16, 45, 20, 32, 8, 18, 78, 62}
	fmt.Println(selectionSort(numbers))
}
