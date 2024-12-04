// Package datastructures contains utility data structures.
//
// The datastructures package can be used to sole Advent of Code puzzles.
package datastructures

// Queue implements a dynamic set with the first-in, first-out (FIFO) policy
// using a linked list. In particular the queue contains pointers to two
// objects:
//   - Head: first element of the queue
//   - Tail: last element of the queue
//
// Data can be deqeueud (removal + return of Head element) or enqueued (new
// data added as Tail element).
type Queue struct {
	Head *Element
	Tail *Element
}

// Element type contains integer (sattelite) data and a pointer to the next
// Element (or NIL) if it is the Tail node. Used as building block for queue.
type Element struct {
	Data int
	Next *Element
}

// Add data as a new Element to the end of the queue.
func (q *Queue) Enqueue(data int) {
	if q.isEmpty() {
		node := &Element{Data: data}
		q.Head = node
		q.Tail = node

	}
	q.Tail.Next = &Element{Data: data, Next: nil}
}

// Remove the first element in the queue, or NIL if Queue is empty, and return
// it.
func (q *Queue) Dequeue() *Element {
	if q.isEmpty() {
		return nil
	}
	dequeuedNode := q.Head
	q.Head = dequeuedNode.Next
	return dequeuedNode
}

func NewEmptyQueue() *Queue {
	return &Queue{Head: nil, Tail: nil}
}

func NewQueue(data ...int) *Queue {
	var q = NewEmptyQueue()
	for _, val := range data {
		q.Enqueue(val)
	}

	return q
}

// isEmpty checks if the Queue is empty (Head and Tail are both NIL)
func (q *Queue) isEmpty() bool {
	return q.Head == q.Tail
}
