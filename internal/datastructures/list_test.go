package datastructures_test

import (
	"testing"

	dst "github.com/lo-b/aoc24/internal/datastructures"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name            string
		initialList     *dst.DoublyLinkedList
		insertElements  []*dst.ListElement
		expectedHeadKey int
		expectedTailKey int
	}{
		{
			name:            "Insert element into empty list",
			initialList:     &dst.DoublyLinkedList{Head: nil, Tail: nil},
			insertElements:  []*dst.ListElement{{Key: 10}},
			expectedHeadKey: 10,
			expectedTailKey: 10,
		},
		{
			name:            "Insert elements into list with Head set",
			initialList:     &dst.DoublyLinkedList{Head: createElement(3), Tail: nil},
			insertElements:  []*dst.ListElement{{Key: 10}, {Key: 7}},
			expectedHeadKey: 3,
			expectedTailKey: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.initialList

			for _, element := range tt.insertElements {
				list.Insert(element)
			}

			if list.Head != nil && list.Head.Key != tt.expectedHeadKey {
				t.Errorf(
					"expected Head to have key %d, got %d",
					tt.expectedHeadKey,
					list.Head.Key,
				)
			}

			if list.Tail != nil && list.Tail.Key != tt.expectedTailKey {
				t.Errorf("expected Tail to have key %d, got %d",
					tt.expectedTailKey,
					list.Tail.Key,
				)
			}
		})
	}
}

func TestLinks(t *testing.T) {
	tests := []struct {
		name    string
		list    *dst.DoublyLinkedList
		inserts []*dst.ListElement
	}{
		{
			name:    "Insert 1, 2, 3, 4",
			list:    &dst.DoublyLinkedList{Head: nil, Tail: nil},
			inserts: []*dst.ListElement{createElement(1), createElement(2), createElement(3), createElement(4)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.list
			insertedElements := tt.inserts

			// insert into list
			for _, element := range tt.inserts {
				list.Insert(element)
			}

			if list.Head != insertedElements[0] {
				t.Errorf("expected first inserted element (key: 1) to be Head")
			}

			if insertedElements[0].Next != insertedElements[1] {
				t.Errorf("expected link 1 ==> 2 to be present")
			}

			if insertedElements[1].Prev != insertedElements[0] {
				t.Errorf("expected link 2 <== 1 to be present")
			}

			if insertedElements[1].Next != insertedElements[2] {
				t.Errorf("expected link 2 ==> 3 to be present")
			}

			if insertedElements[2].Prev != insertedElements[1] {
				t.Errorf("expected link 2 <== 3 to be present")
			}

			if insertedElements[2].Next != insertedElements[3] {
				t.Errorf("expected link 3 ==> 4 to be present")
			}

			if insertedElements[3].Prev != insertedElements[2] {
				t.Errorf("expected link 3 <== 4 to be present")
			}

			if list.Tail != tt.inserts[3] {
				t.Errorf("expected last inserted element (key: 4) to be Tail")
			}

		})
	}
}

func TestDelete_DeleteHeadTail_EmptyListLeft(t *testing.T) {

	list := &dst.DoublyLinkedList{Head: nil, Tail: nil}
	insertElements := []*dst.ListElement{{Key: 1}, {Key: 2}}

	for _, element := range insertElements {
		list.Insert(element)
	}

	list.Delete(list.Head)
	list.Delete(list.Tail)

	if list.Tail != nil {
		t.Error("expected Tail to be NIL")
	}

	if list.Head != nil {
		t.Error("expected Head to be NIL")
	}
}

func TestDelete_DeleteHeadTail_SingleElementLeft(t *testing.T) {
	// expected key left after removal
	expectedKey := 2

	list := &dst.DoublyLinkedList{Head: nil, Tail: nil}
	insertElements := []*dst.ListElement{{Key: 1}, {Key: 2}, {Key: 3}}

	for _, element := range insertElements {
		list.Insert(element)
	}

	list.Delete(list.Head)
	list.Delete(list.Tail)

	if list.Tail != list.Head || list.Head.Key != expectedKey {
		t.Errorf("expected single element with key %d to be present.", expectedKey)
	}
}

func createElement(key int) *dst.ListElement {
	return &dst.ListElement{
		Key:  key,
		Next: nil,
		Prev: nil,
	}
}
