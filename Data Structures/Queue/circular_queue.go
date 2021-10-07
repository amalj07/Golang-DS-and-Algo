package main

import "fmt"

var size = 5

type Queue struct {
	items []int
	front int
	rear  int
}

func (queue *Queue) isFull() bool {
	if queue.front == 0 && queue.rear == size-1 || queue.front == queue.rear+1 {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

func (queue *Queue) isEmpty() bool {
	if queue.front == -1 && queue.rear == -1 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

func (queue *Queue) enqueue(value int) {
	if queue.isFull() {
		return
	}

	if queue.front == -1 {
		queue.front++
	}

	queue.rear = (queue.rear + 1) % size
	queue.items[queue.rear] = value
	return
}

func (queue *Queue) dequeue() {
	if queue.isEmpty() {
		return
	}

	// element := queue.items[queue.front]
	if queue.front == queue.rear {
		queue.front, queue.rear = -1, -1
		return
	}

	queue.front = (queue.front + 1) % size
	return
}

func (queue *Queue) display() {
	if queue.isEmpty() {
		return
	}

	for i := queue.front; i != queue.rear; i = (i + 1) % size {
		fmt.Print(queue.items[i], " ")
	}

	fmt.Print(queue.items[queue.rear])
}

func main() {
	q := Queue{
		items: make([]int, size),
		front: -1,
		rear:  -1,
	}

	q.enqueue(3)
	q.enqueue(8)
	q.enqueue(6)
	q.dequeue()
	q.enqueue(2)
	q.enqueue(9)
	q.dequeue()
	q.dequeue()
	q.enqueue(7)
	q.dequeue()
	q.enqueue(2)
	q.display()
}
