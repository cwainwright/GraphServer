package linkedlist

import (
	"errors"
	"fmt"
)

// --- Node ---

type Node[T any] struct {
	Prev  *Node[T] // previous Node in the list (nil if none)
	Next  *Node[T] // next Node in the list (nil if none)
	Value T        // value of Node (type T)
}

// Link nextNode to currNode
func (currNode *Node[T]) linkNext(newNode *Node[T]) {
	// handle case of currNode being inserted (an existing nextNode)
	/* currNode -> nextNode
	  		 	 ^
		   	  newNode
	*/
	if currNode.Next != nil {
		currNode.Next.Prev = newNode
		newNode.Next = currNode.Next
	}
	currNode.Next = newNode
	newNode.Prev = currNode
}

func (currNode *Node[T]) linkPrev(newNode *Node[T]) {
	// handle case of currNode being inserted (and existing prevNode)
	/* currNode <- nextNode
	  			^
		   	 newNode
	*/
	if currNode.Prev != nil {
		currNode.Prev.Next = newNode
		newNode.Prev = currNode.Prev
	}
	currNode.Prev = newNode
	newNode.Next = currNode
}

func (currNode *Node[T]) unlink() {
	// Relink adjacent nodes to each other
	if currNode.Next != nil {
		currNode.Next.Prev = currNode.Prev
	}
	if currNode.Prev != nil {
		currNode.Prev.Next = currNode.Next
	}
	// Unlink self from adjacent nodes
	currNode.Next, currNode.Prev = nil, nil
}

func (currNode *Node[T]) Print() {
	fmt.Printf("Node: %+v", currNode.Value)
}

// --- Linked List ---

type LinkedList[T any] struct {
	Head *Node[T] // start Node of list (nil if empty)
	Tail *Node[T] // end Node of list (nil if empty)
}

func (list *LinkedList[T]) isEmpty() bool { // returns true if head and tail are nil
	return list.Head == nil && list.Tail == nil
}

func (list *LinkedList[T]) isHealthy() bool { // returns true if list is empty or is consistent
	return list.isEmpty() || !(list.Head == nil || list.Tail == nil)
}

func (list *LinkedList[T]) AppendNode(value T) *T {
	newNode := Node[T]{nil, nil, value}
	if list.isEmpty() { // if list is empty
		list.Head = &newNode
		list.Tail = &newNode
		return &newNode.Value
	} else if !list.isHealthy() { // if list is not healthy
		err := errors.New("failed to append node; list is corrupt")
		fmt.Println("Error:", err)
		return nil
	} else { // otherwise append node
		list.Tail.linkNext(&newNode)
		list.Tail = &newNode
		return &newNode.Value
	}
}

func (list *LinkedList[T]) PrependNode(value T) *T {
	newNode := Node[T]{nil, nil, value}
	if list.isEmpty() { // if list is empty
		list.Head = &newNode
		list.Tail = &newNode
		return &newNode.Value
	} else if !list.isHealthy() { // if list is not healthy
		err := errors.New("failed to prepend node; list is corrupt")
		fmt.Println("Error:", err)
		return nil
	} else { // otherwise prepend node
		list.Head.linkPrev(&newNode)
		list.Head = &newNode
		return &newNode.Value
	}
}

func (list *LinkedList[T]) RemoveNode(eq func(T, T) bool, value T) bool {
	for node := list.Head; node != nil; node = node.Next {
		if eq(node.Value, value) {
			if list.Head == node {
				list.Head = node.Next
			}
			if list.Tail == node {
				list.Tail = node.Prev
			}
			node.unlink()
			return true
		}
	}
	return false
}

func (list *LinkedList[T]) PopNode() *T {
	if list.isEmpty() {
		return nil
	} else {
		node := list.Tail
		list.Tail = node.Prev
		node.unlink()
		return &node.Value
	}
}

func (list *LinkedList[T]) Print() {
	for node := list.Head; node != nil; node = node.Next {
		node.Print()
		fmt.Print(", ")
	}
}
