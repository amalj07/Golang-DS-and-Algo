package main

import "fmt"

type maxHeap struct {
	data []int
}

func (m *maxHeap) heapifyUp(index, parent int) {
	if m.data[index] > m.data[parent] {
		m.data[index], m.data[parent] = m.data[parent], m.data[index]
	}

	if parent != 0 {
		index = parent
		parent = (index - 1) / 2
		m.heapifyUp(index, parent)
	}
}

func (m *maxHeap) heapifyDown(index int) {
	leftChild := 2 * index + 1
	rightChild := 2 * index + 2

	larger := index
	if leftChild <= len(m.data) - 1 && m.data[leftChild] > m.data[larger] {
		larger = leftChild
	}

	if rightChild <= len(m.data) - 1 && m.data[rightChild] > m.data[larger] {
		larger = rightChild
	}

	if larger != index {
		m.data[index], m.data[larger] = m.data[larger], m.data[index]
		m.heapifyDown(larger)
	}
}

func (m *maxHeap) insert(value int) {
	m.data = append(m.data, value)
	index := len(m.data) - 1
	parent := (index - 1) / 2

	m.heapifyUp(index, parent)
}

func (m *maxHeap) extract() int {
	if len(m.data) == 0 {
		fmt.Println("Heap is emtpy")
		return -1
	}
	extractedValue := m.data[0]
	m.data[0] = m.data[len(m.data)-1] //Move last element in heap to the root
	m.data = m.data[:len(m.data)-1]   //Decrease the length of the array

	m.heapifyDown(0)

	return extractedValue
}

func main() {
	m := maxHeap{
		data: []int{80, 69, 73, 56, 45, 51, 52},
	}
	m.insert(70)
	m.extract()
	m.extract()
}
