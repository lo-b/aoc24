package datastructures

type DoubleLinkedList struct {
	Head *ListElement
	Tail *ListElement
}

type ListElement struct {
	Key  int
	Next *ListElement
	Prev *ListElement
}

func (list *DoubleLinkedList) Insert(element *ListElement) {
	if list.Head == nil {
		list.Head = element
	}
	element.Prev = list.Tail
	if list.Tail != nil {
		list.Tail.Next = element
	}
	list.Tail = element
	element.Next = nil
}

func (list *DoubleLinkedList) Delete(element *ListElement) {
	// update link to Next of previous element
	if element.Prev != nil {
		element.Prev.Next = element.Next
		// if deleting the head element, only need to reset the pointer of Head
	} else {
		list.Head = element.Next
	}
	// update link to Prev element
	if element.Next != nil {
		element.Next.Prev = element.Prev
	} else {
		list.Tail = element.Prev
	}
}
