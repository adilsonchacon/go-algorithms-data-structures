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

func (node *ListNode) Add(value int) {
	if node.Next == nil {
		node.Next = &ListNode{Value: value, Previous: node}
	} else if node.Next.Value > value {
		newNode := &ListNode{Value: value, Next: node.Next, Previous: node}
		node.Next.Previous = newNode
		node.Next = newNode
	} else {
		node.Next.Add(value)
	}
}

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

func (list *List) SetTail() *ListNode {
	node := list.Tail
	if node == nil {
		node = list.Head
	}

	if node == nil {
		return nil
	}

	for {
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	return node
}

func (list *List) Add(value int) {
	if list.Head == nil {
		list.Head = &ListNode{Value: value}
	} else if list.Head.Value > value {
		aux := &ListNode{Value: value, Next: list.Head}
		list.Head.Previous = aux
		list.Head = aux
	} else {
		list.Head.Add(value)
	}

	list.Tail = list.SetTail()
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
