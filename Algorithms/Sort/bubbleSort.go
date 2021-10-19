package main

import "fmt"

func bubbleSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	numbers := []int{25, 25, 10, 3, 78, 16, 34, 89, 0, 29, 55}
	fmt.Println(bubbleSort(numbers))
}
