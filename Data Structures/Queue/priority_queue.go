package main

import "fmt"

type priorityQueue struct {
	data []int
}

func (pq *priorityQueue) enQueue(value int) {
	pq.data = append(pq.data, value)
	childIdx := len(pq.data) - 1
	parent := (childIdx - 1) / 2

	for pq.data[childIdx] > pq.data[parent] {
		pq.data[childIdx], pq.data[parent] = pq.data[parent], pq.data[childIdx]
		childIdx = parent
		parent = (childIdx - 1) / 2
	}

	return
}

func (pq *priorityQueue) deQueue() int {
	pq.data[0], pq.data[len(pq.data)-1] = pq.data[len(pq.data)-1], pq.data[0] // Swap root and last element
	value := pq.data[len(pq.data)-1]
	pq.data = pq.data[:len(pq.data)-1] // Decrease the length of data

	pq.heapify(pq.data, len(pq.data), 0)

	return value
}

func (pq *priorityQueue) heapify(queue []int, size, i int) {
	fmt.Println("h", queue, i)
	largest := i
	leftChild := 2*i + 1
	rightChild := 2*i + 2

	fmt.Println(largest, leftChild, rightChild)

	if leftChild <= size && queue[leftChild] > queue[largest] {
		largest = leftChild
	}

	if rightChild <= size && queue[rightChild] > queue[largest] {
		largest = rightChild
	}

	fmt.Println("lar", largest, "i", i)

	if largest != i {
		queue[i], queue[largest] = queue[largest], queue[i]
		pq.heapify(queue, len(queue)-1, largest)
	}
}

func main() {
	pq := priorityQueue{
		data: []int{},
	}

	pq.enQueue(3)
	pq.enQueue(8)
	pq.enQueue(10)
	pq.enQueue(13)
	pq.enQueue(21)
	pq.enQueue(5)
	pq.enQueue(16)
	pq.enQueue(7)
	pq.deQueue()

	fmt.Println(pq.data)
}