package datastructures

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	tests := []struct {
		name         string
		initialQueue *Queue
		enqueues     []int
		expectedTail int
	}{
		{
			name:         "enqueue to empty queue",
			initialQueue: NewEmptyQueue(),
			enqueues:     []int{10},
			expectedTail: 10,
		},
		{
			name:         "enqueue multiple items",
			initialQueue: NewQueue(1, 1),
			enqueues:     []int{2, 3, 4},
			expectedTail: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.initialQueue

			for _, data := range tt.enqueues {
				q.Enqueue(data)
			}

			if q.Tail == nil || q.Tail.Data != tt.expectedTail {
				t.Errorf("expected tail to have data %v, got %v", tt.expectedTail, q.Tail)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	tests := []struct {
		name          string
		initialQueue  *Queue
		dequeueAmount int
		expectedHead  *Element
	}{
		{
			name:          "dequeue from empty queue",
			initialQueue:  NewEmptyQueue(),
			dequeueAmount: 1,
			expectedHead:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.initialQueue

			for i := 0; i < tt.dequeueAmount; i++ {
				q.Dequeue()
			}

			if q.Head != tt.expectedHead {
				t.Errorf("expected head to point to %v, got %v", q.Head, q.Tail)
			}
		})
	}
}
