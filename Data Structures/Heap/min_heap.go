package main

import "fmt"

type minHeap struct {
	data []int
}

func (m *minHeap) heapifyDown(index int) {
	smaller := index
	leftChild := 2 * index + 1
	rightChild := 2 * index + 2

	if leftChild <= len(m.data)-1 && m.data[leftChild] < m.data[smaller] {
		smaller = leftChild
	}

	if rightChild <= len(m.data) - 1 && m.data[rightChild] < m.data[smaller] {
		smaller = rightChild
	}

	if smaller != index {
		m.data[index], m.data[smaller] = m.data[smaller], m.data[index]
		m.heapifyDown(smaller)
	}
}

func (m *minHeap) heapifyUp(index, parent int) {
	if m.data[index] < m.data[parent] {
		m.data[index], m.data[parent] = m.data[parent], m.data[index]
	}
	if parent != 0 {
		index = parent
		parent = (index - 1) / 2
		m.heapifyUp(index, parent)
	}
}

func (m *minHeap) insert(value int) {
	m.data = append(m.data, value)
	index := len(m.data) - 1
	parent := (index - 1) / 2

	m.heapifyUp(index, parent)
}

func (m *minHeap) extract() int {
	if len(m.data) == 0 {
		fmt.Println("Heap is empty")
	}

	m.data[0], m.data[len(m.data)-1] = m.data[len(m.data)-1], m.data[0]
	value := m.data[len(m.data)-1]
	m.data = m.data[:len(m.data)-1]

	m.heapifyDown(0)

	return value
}

func main() {
	m := minHeap{}
	m.insert(2)
	m.insert(7)
	m.insert(0)
	m.insert(4)
	m.insert(9)
	m.insert(3)
	m.insert(6)
	m.extract()

	fmt.Println(m.data)
}
