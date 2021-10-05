package main

import "fmt"

var size = 5

type Queue struct {
	items []int
	front int
	rear  int
}

func (queue *Queue) enqueue(value int) {
	if queue.rear == size-1 {
		fmt.Println("Queue is full")
		return
	}

	queue.items = append(queue.items, value)
	if queue.front == -1 {
		queue.front, queue.rear = 0, 0
	} else {
		queue.rear++
	}
	return
}

func (queue *Queue) dequeue() {
	if queue.front == -1 {
		fmt.Println("Queue is emtpy")
		return
	}

	queue.items = queue.items[1:]
	queue.rear--
	return
}

func (queue *Queue) isEmtpy() {
	if queue.front == -1 {
		fmt.Println("Queue is emtpy")
	} else {
		fmt.Println("Queue is not empty")
	}
}

func (queue *Queue) isFull() {
	if queue.rear == size-1 {
		fmt.Println("Queue is full")
	} else {
		fmt.Println("Queue is no full")
	}
}

func main() {
	q := Queue{
		front: -1,
		rear:  -1,
	}

	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(7)
	q.dequeue()
	q.isFull()
	q.isEmtpy()
}
