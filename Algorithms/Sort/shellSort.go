package main

import "fmt"

func shellSort(arr []int) []int {
	n := len(arr)

	for h := n/2; h >= 1; h=h/2 {
		for  i := h; i < n; i++ {
			for j := i; j >= h && arr[j] < arr[j-h] ; j = j-h{
					arr[j], arr[j-h] = arr[j-h], arr[j]
			}	
		}
	}

	return arr
}

func main() {
	numbers := []int{23, 12, 3, 56, 16, 45, 20, 32, 8, 18, 78, 62}
	fmt.Println(shellSort(numbers))
}