package queues

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// NewNode queue node
func NewNode(data int) *Node {
	return &Node{data, nil}
}

// Queue struct
type Queue struct {
	size int
	head *Node
	tail *Node
}

// NewQueue func
func NewQueue() *Queue {
	return &Queue{0, nil, nil}
}

// Enqueue to queue
func (q *Queue) Enqueue(data int) {
	n := NewNode(data)
	last := &(q.tail)
	first := &(q.head)

	if *last != nil {
		(*last).next = n
	}
	*last = n
	if *first == nil {
		*(first) = *last
	}
	q.size++
}

// Dequeue from queue
func (q *Queue) Dequeue() int {
	first := &(q.head)
	last := &(q.tail)
	if *first == nil {
		fmt.Println("Empty queue")
	} else {
		res := (*first).data
		if (*first).next == nil {
			*last = nil
		}
		*first = (*first).next
		q.size--
		return res
	}

	return -1
}
