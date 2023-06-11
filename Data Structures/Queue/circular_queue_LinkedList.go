package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type queue struct {
	front *Node
	rear  *Node
}

func (q *queue) isEmpty() bool {
	if q.front == nil && q.rear == nil {
		return true
	}
	return false
}

func (q *queue) display() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	ptr := q.front

	for ptr != q.rear {
		fmt.Print(ptr.value)
		ptr = ptr.next
		fmt.Print(" ")
	}
	fmt.Print(ptr.value)
	return
}

func (q *queue) enqueue(value int) {
	newNode := Node{
		value: value,
	}

	if q.front == nil {
		q.front = &newNode
		q.rear = &newNode
		return
	}

	q.rear.next = &newNode
	q.rear = &newNode
	newNode.next = q.front
	return
}

func (q *queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	if q.front == q.rear {
		q.front, q.rear = nil, nil
		return
	}

	q.front = q.front.next
	q.rear.next = q.front
	return
}


func main() {
	q := queue{}
	q.enqueue(3)
	q.enqueue(4)
	q.dequeue()
	q.enqueue(2)
	q.enqueue(1)
	q.dequeue()
	q.enqueue(7)
	q.enqueue(1)
	q.display()
}