package main

import "fmt"

type MaxHeap struct {
	data []int
}

func (maxheap *MaxHeap) insert(value int) {
	maxheap.data = append(maxheap.data, value)
	index := len(maxheap.data) - 1
	parent := (index - 1) / 2
	for maxheap.data[parent] < maxheap.data[index] {
		maxheap.data[parent], maxheap.data[index] = maxheap.data[index], maxheap.data[parent]
		index = parent
		parent = (index - 1) / 2
	}
	return
}

func (maxheap *MaxHeap) extract() int {
	if len(maxheap.data) == 0 {
		fmt.Println("Heap is emtpy")
		return -1
	}
	extractedValue := maxheap.data[0]
	maxheap.data[0] = maxheap.data[len(maxheap.data)-1] //Move last element in heap to the root
	maxheap.data = maxheap.data[:len(maxheap.data)-1]   //Decrease the length of the array
	parentIndex, leftChildIndex, rightChildIndex := 0, 1, 2

	// Loop until the last left child
	for leftChildIndex <= len(maxheap.data)-1 {
		if leftChildIndex == len(maxheap.data)-1 { // Left child is the last child(parent has only one child)
			if maxheap.data[parentIndex] < maxheap.data[leftChildIndex] {
				// Parent less than left child
				maxheap.data[parentIndex], maxheap.data[leftChildIndex] = maxheap.data[leftChildIndex], maxheap.data[parentIndex]
				return extractedValue
			}
			return extractedValue
		} else if maxheap.data[leftChildIndex] > maxheap.data[rightChildIndex] {
			if maxheap.data[parentIndex] < maxheap.data[leftChildIndex] {
				// Left child is greater than right child and parent is less than right child(swap parent and left child)
				maxheap.data[parentIndex], maxheap.data[leftChildIndex] = maxheap.data[leftChildIndex], maxheap.data[parentIndex]
			}
			parentIndex = leftChildIndex
			leftChildIndex = 2*parentIndex + 1
			rightChildIndex = 2*parentIndex + 2
		} else if maxheap.data[leftChildIndex] < maxheap.data[rightChildIndex] {
			if maxheap.data[parentIndex] < maxheap.data[rightChildIndex] {
				// Left child is less than right child and parent is less than right child(swap parent and right child)
				maxheap.data[parentIndex], maxheap.data[rightChildIndex] = maxheap.data[rightChildIndex], maxheap.data[parentIndex]
			}
			parentIndex = rightChildIndex
			leftChildIndex = 2*parentIndex + 1
			rightChildIndex = 2*parentIndex + 2
		} else { // Both child are equal
			if maxheap.data[parentIndex] < maxheap.data[leftChildIndex] {
				// Parent less than left child(swap parent and left child)
				maxheap.data[parentIndex], maxheap.data[leftChildIndex] = maxheap.data[leftChildIndex], maxheap.data[parentIndex]
				parentIndex = leftChildIndex
				leftChildIndex = 2*parentIndex + 1
				rightChildIndex = 2*parentIndex + 2
				continue
			}
			return extractedValue
		}
	}

	return extractedValue
}

func main() {
	h := MaxHeap{
		data: []int{80, 69, 73, 56, 45, 51, 52},
	}
	h.insert(70)
	fmt.Println(h.data)
	h.extract()
	fmt.Println(h.data)
	h.extract()
	fmt.Println(h.data)
}
