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

	if queue.front == -1 {
		queue.front = 0
	}
	queue.rear++
	queue.items[queue.rear] = value
	return
}

func (queue *Queue) dequeue() {
	if queue.front == -1 {
		fmt.Println("Queue is emtpy")
		return
	}

	queue.front++
	if queue.front > queue.rear {
		queue.front, queue.rear = -1, -1
	}
	return
}

func (queue *Queue) display() {
	if queue.front == -1 {
		fmt.Println("Queue is emtpy")
		return
	}

	for i := queue.front; i < queue.rear; i++ {
		fmt.Print(queue.items[i], " ")
	}
	fmt.Print(queue.items[queue.rear])
}

func (queue *Queue) peek() {
	if queue.front == -1 {
		fmt.Println("Queue is emtpy")
		return
	}

	fmt.Println(queue.items[queue.front])
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
		fmt.Println("Queue is not full")
	}
}

func main() {
	q := Queue{
		items: make([]int, size),
		front: -1,
		rear:  -1,
	}

	q.enqueue(3)
	q.enqueue(2)
	q.enqueue(7)
	q.dequeue()
	q.enqueue(6)
	q.enqueue(1)
	q.dequeue()
	q.peek()
	q.display()
}
