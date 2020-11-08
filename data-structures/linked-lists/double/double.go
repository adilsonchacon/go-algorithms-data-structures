package main

import (
	"fmt"
)

type ListNode struct {
	Value    int
	Next     *ListNode
	Previous *ListNode
}

type List struct {
	Head *ListNode
	Tail *ListNode
}

// ListNode

func (node *ListNode) PrintForward() {
	if node != nil {
		fmt.Println(node.Value)
		node.Next.PrintForward()
	}
}

func (node *ListNode) PrintBackward() {
	if node != nil {
		fmt.Println(node.Value)
		node.Previous.PrintBackward()
	}
}

// List

func (list *List) Add(value int) {
	newNode := &ListNode{Value: value}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = list.Head
	} else {
		newNode.Previous = list.Tail
		list.Tail.Next = newNode
		list.Tail = list.Tail.Next
	}
}

func (list *List) PrintForward() {
	if list.Head == nil {
		fmt.Println("List is empty")
	} else {
		list.Head.PrintForward()
	}
}

func (list *List) PrintBackward() {
	if list.Tail == nil {
		fmt.Println("List is empty")
	} else {
		list.Tail.PrintBackward()
	}
}

func main() {
	list := &List{}
	list.Add(2)
	list.Add(6)
	list.Add(7)
	list.Add(4)
	list.Add(1)

	fmt.Println("Forward:")
	list.PrintForward()

	fmt.Println("Backward:")
	list.PrintBackward()
}
