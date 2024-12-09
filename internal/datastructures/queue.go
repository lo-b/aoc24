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
	Head *QueueElement
	Tail *QueueElement
}

// QueueElement type contains integer Key (sattelite data) and a pointer to the
// Next QueueElement (or NIL) if it is the Tail node. Used as a building block
// for a Queue.
type QueueElement struct {
	Key  int
	Next *QueueElement
}

// Enqueue adds data as a new Element to the end of the queue.
func (q *Queue) Enqueue(data int) {
	newNode := &QueueElement{Key: data}

	if q.isEmpty() {
		// Queue is empty; Head and Tail will point to the new node
		q.Head = newNode
		q.Tail = newNode
	} else {
		// Append to the end and move the Tail pointer
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

// Dequeue removes the first element in the queue, or NIL if Queue is empty, and return
// it.
func (q *Queue) Dequeue() *QueueElement {
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

// NewQueue creates a new Queue using variable int arg(s).
func NewQueue(data ...int) *Queue {
	var q = NewEmptyQueue()
	for _, val := range data {
		q.Enqueue(val)
	}

	return q
}

// NewQueueFromArray creates new Queue from an array of integers.
func NewQueueFromArray(data []int) *Queue {
	var q = NewEmptyQueue()
	for _, val := range data {
		q.Enqueue(val)
	}

	return q
}

// isEmpty checks if the Queue is empty (i.e. Tail is NIL)
func (q *Queue) isEmpty() bool {
	return q.Tail == nil
}
